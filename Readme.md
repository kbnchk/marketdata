## MARKETDATA is a set of simple tools for getting market data for some cryptocurrency trading platforms

Currently it supports:
- garantex.io,
- beribit.com,
- binance
- binance p2p.
Library is in early development.

### Usage examples:

```go
data, err := marketdata.GarantexNew().GetDOM("usdtrub")
if err != nil {
    panic(err)
}
```
```go
var params url.Values
	params.Add("market", "usdtbtc")
	params.Add("limit", "1000")
data, err := marketdata.GarantexNew().GetHistory(params)
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
		PayTypes:  []string{},
	}
data, err := BinanceP2PNew().GetDOM(config)
if err != nil {
    panic(err)
}
```
