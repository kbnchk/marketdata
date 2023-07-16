package marketdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type garantex struct {
	baseurl, api, dom, history string
}

func Garantex() garantex {
	return garantex{
		baseurl: "https://garantex.io",
		api:     "api/v2",
		dom:     "/depth",
		history: "/trades",
	}
}

//####################################################################
// Depth Of Market
//####################################################################

// GetDOM recieves depth of market
//
// Availible markets:  btcrub, usdtrub, dairub, ethrub, usdcrub, btcusdt, ethbtc, ethusdt, usdcusdt
func (g garantex) GetDOM(market string) (DOM, error) {
	u, _ := url.ParseRequestURI(g.baseurl)
	u.Path = g.api + g.dom
	params := url.Values{"market": []string{market}}
	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return DOM{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return DOM{}, fmt.Errorf("request returned bad status code %s", resp.Status)
	}
	respBody, _ := io.ReadAll(resp.Body)
	respBytes := []byte(respBody)

	var model garantexDOMResponse
	err = json.Unmarshal(respBytes, &model)
	if err != nil {
		return DOM{}, err
	}
	result := model.toEntity()
	return result, nil
}

// DOM Response
type garantexDOMResponse struct {
	Timestamp float64               `json:"timestamp"`
	Bids      []garantexDOMPosition `json:"bids"`
	Asks      []garantexDOMPosition `json:"asks"`
}
type garantexDOMPosition struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
	Factor string `json:"factor"`
}

func (b garantexDOMResponse) toEntity() DOM {
	bids := make([]DOMPosition, 0, len(b.Bids))
	asks := make([]DOMPosition, 0, len(b.Asks))
	for _, p := range b.Bids {
		bids = append(bids, p.toEntity())
	}
	for _, p := range b.Asks {
		asks = append(asks, p.toEntity())
	}
	return DOM{
		Date: time.Unix(int64(b.Timestamp), 0),
		Bids: bids,
		Asks: asks,
	}
}

func (p garantexDOMPosition) toEntity() DOMPosition {
	return DOMPosition{
		Price:  p.Price,
		Amount: p.Amount,
	}
}

// History

type GarantexHistoryConfig struct {
	Market string // btcrub, usdtrub, dairub, ethrub, usdcrub, btcusdt, ethbtc, ethusdt, usdcusdt
	Limit  int    // optional records amount (default 50, max 1000)
	From   int    // optional trade ID to get data from (but not including)
	To     int    // optional trade ID to get data to (but not including)
	Order  string // optional sorting order ASC DESC
}

func (c GarantexHistoryConfig) toParams() url.Values {
	values := make(url.Values)
	values.Add("market", c.Market)
	if c.Limit != 0 {
		values.Add("limit", strconv.Itoa(c.Limit))
	}
	if c.From != 0 {
		values.Add("from", strconv.Itoa(c.From))
	}
	if c.To != 0 {
		values.Add("to", strconv.Itoa(c.To))
	}
	if c.Order != "" {
		values.Add("order_by", c.Order)
	}
	return values
}

// GetHistory recieves market trading history.
func (g garantex) GetHistory(config GarantexHistoryConfig) ([]HistoryPosition, error) {
	raw, err := g.getHistory(config.toParams())
	if err != nil {
		return nil, err
	}
	return raw.toEntity(), nil

}

func (g garantex) getHistory(params url.Values) (garantexHistoryResponce, error) {
	u, _ := url.ParseRequestURI(g.baseurl)
	u.Path = g.api + g.history
	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request returned bad status code %s", resp.Status)
	}

	respBody, _ := io.ReadAll(resp.Body)
	respBytes := []byte(respBody)

	var model garantexHistoryResponce
	err = json.Unmarshal(respBytes, &model)
	return model, err
}

// History response

type garantexHistoryResponce []garantexHistoryResponcePosition

func (r garantexHistoryResponce) toEntity() []HistoryPosition {
	data := make([]HistoryPosition, 0, len(r))
	for _, el := range r {
		data = append(data, el.toEntity())
	}
	return data
}

// History responce position

type garantexHistoryResponcePosition struct {
	ID     float64 `json:"id"`
	Date   string  `json:"created_at"`
	Price  string  `json:"price"`
	Volume string  `json:"volume"`
	Funds  string  `json:"funds"`
}

func (p garantexHistoryResponcePosition) toEntity() HistoryPosition {
	date, _ := time.Parse(time.RFC3339, p.Date)
	return HistoryPosition{
		ID:     uint(p.ID),
		Date:   date,
		Price:  p.Price,
		Volume: p.Volume,
	}
}
