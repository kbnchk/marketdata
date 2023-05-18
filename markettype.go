package marketdata

import "strings"

type MarketType int

const (
	USDTRUB MarketType = iota
	BTCRUB
	USDCRUB
	ETHRUB
	DAIRUB
)

func (m MarketType) string() string {
	switch m {
	case USDTRUB:
		return "USDT/RUB"
	case BTCRUB:
		return "BTC/RUB"
	case USDCRUB:
		return "USDC/RUB"
	case ETHRUB:
		return "ETH/RUB"
	case DAIRUB:
		return "DAI/RUB"
	default:
		return "unknown"
	}
}

func (m MarketType) name() string {
	return strings.ToLower(strings.Replace(m.string(), "/", "", 1))
}
