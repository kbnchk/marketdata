package marketdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func GetGarantexData(m Market) (Response, error) {
	if m.String() == "unknown" {
		return Response{}, fmt.Errorf("unknown market type")
	}
	url := m.garantexUrl()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Response{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	respBytes := []byte(respBody)

	var model garantexResponse
	err = json.Unmarshal(respBytes, &model)
	if err != nil {
		return Response{}, err
	}
	result := model.toEntity(m)
	return result, nil
}

type garantexResponse struct {
	Timestamp float64            `json:"timestamp"`
	Bids      []garantexPosition `json:"bids"`
	Asks      []garantexPosition `json:"asks"`
}

func (b garantexResponse) toEntity(m Market) Response {
	var bids, asks []Position
	for _, p := range b.Bids {
		bids = append(bids, p.convert())
	}
	for _, p := range b.Asks {
		asks = append(asks, p.convert())
	}
	return Response{
		MarketPlace: "garantex",
		MarketName:  m.String(),
		Date:        time.Unix(int64(b.Timestamp), 0),
		Bids:        bids,
		Asks:        asks,
	}
}

type garantexPosition struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
	Factor string `json:"factor"`
}

func (p garantexPosition) convert() Position {
	price, _ := strconv.ParseFloat(p.Price, 64)
	amount, _ := strconv.ParseFloat(p.Amount, 64)
	factor, _ := strconv.ParseFloat(p.Factor, 64)
	var pt string
	switch p.Type {
	case "limit":
		pt = "Фиксированная цена"
		// otherTypes
	default:
		pt = p.Type
	}
	return Position{
		Price:  price,
		Amount: amount,
		Type:   pt,
		Factor: factor,
	}
}
