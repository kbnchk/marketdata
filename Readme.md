## MARKETDATA is a set of simple tools for getting market data for some cryptocurrency trading platforms

Currently it supports:
- garantex.io,
- beribit.com,
- binance
- binance p2p.

Library is in early development.

### Usage examples:

```go
data, err := marketdata.Garantex().GetDOM("usdtrub")
if err != nil {
    panic(err)
}
```
```go
config := GarantexHistoryConfig{
	Market: "usdtbtc",
	Limit:  1000,
	From:   343434,
}

data, err := marketdata.Garantex().GetHistory(config)
if err != nil {
    panic(err)
}
```
```go
config := BinanceP2PConfig{
		Fiat:      "USD",
		Asset:     "USDT",
		Page:      1,
		Rows:      10,
		Countries: []string{"AE"},
	}
data, err := BinanceP2P().GetDOM(config)
if err != nil {
    panic(err)
}
```
