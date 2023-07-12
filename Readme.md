## MARKETDATA is a simple go library for getting market data for some cryptocurrency trading platforms

Currently it supports some data from garantex.io, beribit.com, binance and binance p2p.
It uses garantex.io API and beribit.com websocket.
Library is in early development.

### Usage example:

```go
data, err := marketdata.GarantexNew().GetDOM("usdtrub")
if err != nil {
    panic(err)
}
```