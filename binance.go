package marketdata

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type Binance struct {
	domURL string
	//historyURL string
}

func BinanceNew() Binance {
	return Binance{
		domURL: "https://api.binance.com/api/v3/depth",
	}
}

//####################################################################
// Depth Of Market
//####################################################################

type binanceDOMResponse struct {
	ID   float64    `json:"lastUpdateId"`
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

func (r binanceDOMResponse) toEntity(market string) DOM {
	bids := make([]DOMPosition, 0, len(r.Bids))
	asks := make([]DOMPosition, 0, len(r.Asks))
	for _, p := range r.Bids {
		bids = append(bids, DOMPosition{
			Price:  p[0],
			Amount: p[1],
		})
	}
	for _, p := range r.Asks {
		asks = append(asks, DOMPosition{
			Price:  p[0],
			Amount: p[1],
		})
	}
	return DOM{
		MarketPlace: "binance",
		MarketName:  strings.ToUpper(market),
		Date:        time.Now().UTC(),
		Bids:        bids,
		Asks:        asks,
	}
}

func (g Binance) GetDOM(market string) (DOM, error) {
	url := g.domURL + "?symbol=" + strings.ToUpper(market)
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

	var model binanceDOMResponse
	err = json.Unmarshal(respBytes, &model)
	if err != nil {
		return DOM{}, err
	}
	result := model.toEntity(market)
	return result, nil
}
