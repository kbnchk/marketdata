package marketdata

import "strings"

type MarketType int

const (
	USDTRUB MarketType = iota
	USDTTHB
	USDTRTY
)

func (m MarketType) string() string {
	switch m {
	case USDTRUB:
		return "USDT/RUB"
	case USDTTHB:
		return "USDT/THB"
	case USDTRTY:
		return "USDT/TRY"
	default:
		return "unknown"
	}
}

func (m MarketType) name() string {
	return strings.ToLower(strings.Replace(m.string(), "/", "", 1))
}
