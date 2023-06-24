package marketdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type garantex struct {
	apiBaseURL string
}

func (g garantex) GetDOM(m MarketType) (DOM, error) {
	if m.string() == "unknown" {
		return DOM{}, fmt.Errorf("unknown market type for Garantex marketplace")
	}
	url := g.apiBaseURL + m.name()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DOM{}, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return DOM{}, err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	respBytes := []byte(respBody)

	var model garantexResponse
	err = json.Unmarshal(respBytes, &model)
	if err != nil {
		return DOM{}, err
	}
	result := model.toEntity(m)
	return result, nil
}

type garantexResponse struct {
	Timestamp float64            `json:"timestamp"`
	Bids      []garantexPosition `json:"bids"`
	Asks      []garantexPosition `json:"asks"`
}

func (b garantexResponse) toEntity(m MarketType) DOM {
	bids := make([]DOMPosition, 0, len(b.Bids))
	asks := make([]DOMPosition, 0, len(b.Asks))
	for _, p := range b.Bids {
		bids = append(bids, p.convert())
	}
	for _, p := range b.Asks {
		asks = append(asks, p.convert())
	}
	return DOM{
		MarketPlace: "garantex",
		MarketName:  m.string(),
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

func (p garantexPosition) convert() DOMPosition {
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
	return DOMPosition{
		Price:  price,
		Amount: amount,
		Type:   pt,
		Factor: factor,
	}
}
