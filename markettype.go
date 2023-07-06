package marketdata

import "strings"

type MarketType int

const (
	USDTRUB MarketType = iota
	USDTTHB
	USDTTRY
)

func (m MarketType) string() string {
	switch m {
	case USDTRUB:
		return "USDT/RUB"
	case USDTTHB:
		return "USDT/THB"
	case USDTTRY:
		return "USDT/TRY"
	default:
		return "unknown"
	}
}

func (m MarketType) base() string {
	return strings.Split(m.string(), "/")[0]
}

func (m MarketType) quote() string {
	return strings.Split(m.string(), "/")[1]
}

func (m MarketType) name() string {
	return strings.ToLower(strings.Replace(m.string(), "/", "", 1))
}
