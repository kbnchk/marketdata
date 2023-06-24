package marketdata

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

type beribit struct {
	apiBaseURL string
}

func (b beribit) GetDOM(m MarketType) (DOM, error) {
	if m.string() == "unknown" {
		return DOM{}, fmt.Errorf("unknown market type for Beribit marketplace")
	}
	var resp beribitResponse
	origin := "http://localhost/"
	server := b.apiBaseURL + m.name()
	conf, err := websocket.NewConfig(server, origin)
	if err != nil {
		return DOM{}, err
	}
	ws, err := websocket.DialConfig(conf)
	if err != nil {
		return DOM{}, err
	}
	fr, err := ws.NewFrameReader()
	if err != nil {
		return DOM{}, err
	}
	r := bufio.NewReader(fr)
	data, err := io.ReadAll(r)
	if err != nil {
		return DOM{}, err
	}
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return DOM{}, err
	}
	result := resp.toEntity(m)
	return result, nil
}

type beribitResponse struct {
	Timestamp float64           `json:"Timestamp"`
	Bids      []beribitPosition `json:"Bids"`
	Asks      []beribitPosition `json:"Asks"`
}

func (b beribitResponse) toEntity(m MarketType) DOM {
	bids := make([]DOMPosition, 0, len(b.Bids))
	asks := make([]DOMPosition, 0, len(b.Asks))
	for _, p := range b.Bids {
		bids = append(bids, p.convert())
	}
	for _, p := range b.Asks {
		asks = append(asks, p.convert())
	}
	return DOM{
		MarketPlace: "beribit",
		MarketName:  m.string(),
		Date:        time.Unix(int64(b.Timestamp), 0),
		Bids:        bids,
		Asks:        asks,
	}
}

type beribitPosition struct {
	ExchangeRate float64 `json:"ExchangeRate"`
	Size         float64 `json:"Size"`
	Price        float64 `json:"Price"`
	TypeData     string  `json:"TypeData"`
	Factor       float64 `json:"Factor"`
}

func (p beribitPosition) convert() DOMPosition {
	var factor float64
	var pt string
	var rawtype string
	tmp := strings.Split(p.TypeData, " ")
	if len(tmp) > 1 {
		rawtype = tmp[0]
		fstr := strings.Replace(tmp[1], "%", "", -1)
		factor, _ = strconv.ParseFloat(fstr, 64)
	} else {
		rawtype = p.TypeData
	}
	switch rawtype {
	case "ФЦ":
		pt = "Фиксированная цена"
	// other types
	default:
		pt = rawtype
	}
	return DOMPosition{
		Price:  p.ExchangeRate,
		Amount: p.Size,
		Type:   pt,
		Factor: factor,
	}
}
