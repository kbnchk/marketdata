package marketdata

import "fmt"

type MarketDataSource interface {
	GetDOM(MarketType) (DOM, error)
}

func NewSource(m MarketPlace) (MarketDataSource, error) {
	switch m {
	case Garantex:
		return garantex{
			apiBaseURL: "https://garantex.io/api/v2/depth?market=",
		}, nil
	case Beribit:
		return beribit{
			apiBaseURL: "wss://beribit.com/ws/depth/",
		}, nil
	default:
		return nil, fmt.Errorf("unknown MarketPlace")
	}
}
