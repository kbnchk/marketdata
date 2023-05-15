package marketdata

import "time"

type Response struct {
	MarketPlace string
	MarketName  string
	Date        time.Time
	Bids        []Position
	Asks        []Position
}

type Position struct {
	Price  float64
	Amount float64
	Type   string
	Factor float64
}
