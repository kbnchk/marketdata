package marketdata

import "time"

type DOM struct {
	MarketPlace string
	MarketName  string
	Date        time.Time
	Bids        []DOMPosition
	Asks        []DOMPosition
}

type DOMPosition struct {
	Price  float64
	Amount float64
	Type   string
	Factor float64
}
