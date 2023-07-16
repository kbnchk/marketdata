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

type BinanceP2PConfig struct {
	Fiat      string   // fiat currency ticker
	Asset     string   // asset curency ticker
	Page      int      // amount of pages
	Rows      int      // amount of rows on each page
	Countries []string // Countries for fiat currency
	PayTypes  []string // Payment types

}

func BinanceP2PNew() BinanceP2P {
	return BinanceP2P{
		domURL: "https://p2p.binance.com/bapi/c2c/v2/friendly/c2c/adv/search",
	}
}

//####################################################################
// Depth Of Market
//####################################################################

func (g BinanceP2P) GetDOM(config BinanceP2PConfig) (DOM, error) {
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
		"fiat":              strings.ToUpper(config.Fiat),
		"asset":             strings.ToUpper(config.Asset),
		"page":              config.Page,
		"rows":              config.Rows,
		"tradeType":         "SELL",
		"payTypes":          config.PayTypes,
		"countries":         config.Countries,
		"proMerchantAds":    false,
		"shieldMerchantAds": false,
		"publisherType":     nil,
	})
	asksPostBody, _ := json.Marshal(map[string]any{
		"fiat":              strings.ToUpper(config.Fiat),
		"asset":             strings.ToUpper(config.Asset),
		"page":              config.Page,
		"rows":              config.Rows,
		"tradeType":         "BUY",
		"payTypes":          config.PayTypes,
		"countries":         config.Countries,
		"proMerchantAds":    false,
		"shieldMerchantAds": false,
		"publisherType":     nil,
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
		Date: time.Now().UTC(),
		Bids: bids,
		Asks: asks,
	}, nil
}

// binanceP2PDOM responce
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
