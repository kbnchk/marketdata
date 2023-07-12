package marketdata

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"
)

type BinanceP2P struct {
	domURL string
	//historyURL string
}

func BinanceP2PNew() BinanceP2P {
	return BinanceP2P{
		domURL: "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search",
	}
}

//####################################################################
// Depth Of Market
//####################################################################

type binanceP2PDOMResponse struct {
	Data []struct {
		Adv struct {
			Price  string `json:"price"`
			Amount string `json:"tradableQuantity"`
		} `json:"adv"`
	} `json:"data"`
}

func (r binanceP2PDOMResponse) toPositions() []DOMPosition {
	result := make([]DOMPosition, 0, len(r.Data))
	for _, p := range r.Data {
		result = append(result, DOMPosition(p.Adv))
	}
	return result
}

func (g BinanceP2P) GetDOM(fiat, asset string) (DOM, error) {
	get := func(body []byte) ([]DOMPosition, error) {
		responseBody := bytes.NewBuffer(body)
		req, err := http.NewRequest("POST", g.domURL, responseBody)
		req.Header.Add("Content-Type", "application/json")
		if err != nil {
			return nil, err
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		respBytes := []byte(respBody)

		var model binanceP2PDOMResponse
		err = json.Unmarshal(respBytes, &model)
		if err != nil {
			return nil, err
		}

		return model.toPositions(), err
	}

	bidsPostBody, _ := json.Marshal(map[string]any{
		"fiat":      strings.ToUpper(fiat),
		"page":      1,
		"rows":      10,
		"tradeType": "SELL",
		"asset":     strings.ToUpper(asset),
	})
	asksPostBody, _ := json.Marshal(map[string]any{
		"fiat":      strings.ToUpper(fiat),
		"page":      1,
		"rows":      10,
		"tradeType": "BUY",
		"asset":     strings.ToUpper(asset),
	})

	bids, err := get(bidsPostBody)
	if err != nil {
		return DOM{}, err
	}
	asks, err := get(asksPostBody)
	if err != nil {
		return DOM{}, err
	}
	return DOM{
		MarketPlace: "Binance P2P",
		MarketName:  strings.ToUpper(asset + fiat),
		Date:        time.Now().UTC(),
		Bids:        bids,
		Asks:        asks,
	}, nil
}
