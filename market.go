package marketdata

import "strings"

type Market int

const (
	USDTRUB Market = iota
	BTCRUB
	USDCRUB
	ETHRUB
	DAIRUB
)

func (m Market) String() string {
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

func (m Market) garantexUrl() string {
	baseurl := "https://garantex.io/api/v2/depth?market="
	if m.String() == "unknown" {
		return m.String()
	}
	market := strings.ToLower(strings.Replace(m.String(), "/", "", 1))
	return baseurl + market
}

func (m Market) beribitUrl() string {
	switch m {
	case USDTRUB:
		return "wss://beribit.com/ws/depth/usdtrub"
	default:
		return "unknown"
	}
}
