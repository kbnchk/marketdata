package marketdata

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Garantex struct {
	domURL     string
	historyURL string
}

func GarantexNew() Garantex {
	return Garantex{
		domURL:     "https://garantex.io/api/v2/depth",
		historyURL: "https://garantex.io/api/v2/trades",
	}
}

//####################################################################
// Depth Of Market
//####################################################################

func (g Garantex) GetDOM(market string) (DOM, error) {
	url := g.domURL + "?market=" + strings.ToLower(market)
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

	var model garantexDOMResponse
	err = json.Unmarshal(respBytes, &model)
	if err != nil {
		return DOM{}, err
	}
	result := model.toEntity(market)
	return result, nil
}

// DOM Response
type garantexDOMResponse struct {
	Timestamp float64               `json:"timestamp"`
	Bids      []garantexDOMPosition `json:"bids"`
	Asks      []garantexDOMPosition `json:"asks"`
}

func (b garantexDOMResponse) toEntity(market string) DOM {
	bids := make([]DOMPosition, 0, len(b.Bids))
	asks := make([]DOMPosition, 0, len(b.Asks))
	for _, p := range b.Bids {
		bids = append(bids, p.toEntity())
	}
	for _, p := range b.Asks {
		asks = append(asks, p.toEntity())
	}
	return DOM{
		MarketPlace: "garantex",
		MarketName:  strings.ToUpper(market),
		Date:        time.Unix(int64(b.Timestamp), 0),
		Bids:        bids,
		Asks:        asks,
	}
}

type garantexDOMPosition struct {
	Price  string `json:"price"`
	Volume string `json:"volume"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
	Factor string `json:"factor"`
}

func (p garantexDOMPosition) toEntity() DOMPosition {
	return DOMPosition{
		Price:  p.Price,
		Amount: p.Amount,
	}
}

//####################################################################
// History
//####################################################################

func (g Garantex) GetHistoryToDate(market string, earliest time.Time) ([]HistoryPosition, error) {
	const limit = 1000
	data := make([]garantexHistoryPosition, 0, limit)

	//recursively gets positions until reaches earliest parameter
	var getTo func(uint, *[]garantexHistoryPosition) error
	getTo = func(toID uint, data *[]garantexHistoryPosition) error {
		var idstring string

		if toID != 0 {
			idstring = fmt.Sprintf("&to=%d", toID)
		}

		url := g.historyURL + "?market=" + strings.ToLower(market) + fmt.Sprintf("&limit=%d", limit) + idstring
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		respBytes := []byte(respBody)

		var model []garantexHistoryPosition
		err = json.Unmarshal(respBytes, &model)
		if err != nil {
			return err
		}
		if len(model) > 0 {
			*data = append(*data, model...)
			firstentity := model[len(model)-1].toEntity()
			if firstentity.Date.After(earliest) {
				err := getTo(firstentity.ID, data)
				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	err := getTo(0, &data)
	if err != nil {
		return nil, err
	}

	result := make([]HistoryPosition, 0, len(data))
	for _, el := range data {
		position := el.toEntity()
		if position.Date.After(earliest) {
			result = append(result, position)
		} else {
			break
		}
	}
	return result, nil
}

func (g Garantex) GetHistoryFromID(market string, id uint) ([]HistoryPosition, error) {

	const limit = 1000
	data := make([]garantexHistoryPosition, 0, limit)

	var getFrom func(uint, *[]garantexHistoryPosition) error
	getFrom = func(fromID uint, data *[]garantexHistoryPosition) error {

		url := g.historyURL + "?market=" + strings.ToLower(market) + fmt.Sprintf("&limit=%d", limit) + fmt.Sprintf("&order_by=asc&from=%d", fromID)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		respBody, _ := io.ReadAll(resp.Body)
		respBytes := []byte(respBody)

		var model []garantexHistoryPosition
		err = json.Unmarshal(respBytes, &model)
		if err != nil {
			return err
		}

		if len(model) > 0 {
			*data = append(*data, model...)
			if len(model) == limit {
				err := getFrom(uint(model[len(model)-1].ID), data)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	err := getFrom(id, &data)
	if err != nil {
		return nil, err
	}

	result := make([]HistoryPosition, 0, len(data))
	for _, el := range data {
		result = append(result, el.toEntity())
	}
	return result, nil
}

// History response
type garantexHistoryPosition struct {
	ID     float64 `json:"id"`
	Date   string  `json:"created_at"`
	Price  string  `json:"price"`
	Volume string  `json:"volume"`
	Funds  string  `json:"funds"`
}

func (p garantexHistoryPosition) toEntity() HistoryPosition {
	date, _ := time.Parse(time.RFC3339, p.Date)
	return HistoryPosition{
		ID:     uint(p.ID),
		Date:   date,
		Price:  p.Price,
		Volume: p.Volume,
		Funds:  p.Funds,
	}
}
