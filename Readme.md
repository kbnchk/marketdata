## MARKETDATA is a simple go library for getting market data for some cryptocurrency trading platforms

Currently only garantex.io and beribit.com is supported.
It uses garantex.io API and beribit.com websocket.
Library is in early development.

### Usage example:

```go
garantex := marketdata.GarantexNew()
data, err := garantex.GetDOM(marketdata.USDTRUB)
if err != nil {
    panic(err)
}
```