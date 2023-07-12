package marketdata

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

type Beribit struct {
	domURL string
	//historyURL string
}

func BeribitNew() Beribit {
	return Beribit{
		domURL: "wss://beribit.com/ws/depth/",
	}
}

func (b Beribit) GetDOM(market string) (DOM, error) {

	var resp beribitResponse
	origin := "https://beribit.com/"
	server := b.domURL + strings.ToUpper(market)
	conf, err := websocket.NewConfig(server, origin)
	if err != nil {
		return DOM{}, err
	}
	conf.Header = http.Header{"User-Agent": []string{"marketdata-app"}}
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
	result := resp.toEntity(market)
	return result, nil
}

type beribitResponse struct {
	Depth struct {
		MarketName string            `json:"MarketName"`
		Timestamp  float64           `json:"Timestamp"`
		Bids       []beribitPosition `json:"Bids"`
		Asks       []beribitPosition `json:"Asks"`
	} `json:"Depth"`
}

func (b beribitResponse) toEntity(market string) DOM {
	depth := b.Depth
	bids := make([]DOMPosition, 0, len(depth.Bids))
	asks := make([]DOMPosition, 0, len(depth.Asks))
	for _, p := range depth.Bids {
		bids = append(bids, p.toEntity())
	}
	for _, p := range depth.Asks {
		asks = append(asks, p.toEntity())
	}
	return DOM{
		MarketPlace: "beribit",
		MarketName:  strings.ToUpper(market),
		Date:        time.Unix(int64(depth.Timestamp), 0),
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

func (p beribitPosition) toEntity() DOMPosition {
	return DOMPosition{
		Price:  fmt.Sprintf("%g", p.ExchangeRate),
		Amount: fmt.Sprintf("%g", p.Size),
	}
}
