package marketdata

import (
	"encoding/json"
	"fmt"
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

func (r binanceDOMResponse) toEntity(m MarketType) DOM {
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
		MarketName:  m.string(),
		Date:        time.Now().UTC(),
		Bids:        bids,
		Asks:        asks,
	}
}

func (g Binance) GetDOM(m MarketType) (DOM, error) {
	if m.string() == "unknown" {
		return DOM{}, fmt.Errorf("unknown market type")
	}
	url := g.domURL + "?symbol=" + strings.ToUpper(m.name())
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
	result := model.toEntity(m)
	return result, nil
}
