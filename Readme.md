## MARKETDATA is a simple go library for getting market data for some cryptocurrency trading platforms

Currently only DOM (depth of market) fetching is supported, and only for garantex.io and beribit.com.
It uses garantex.io API and beribit.com websocket

### Usage example:
```go
source, err := marketdata.NewSource(Garantex)
if err != nil {
    panic(err)
}
data, err := source.GetDOM(USDTRUB)
if err != nil {
    panic(err)
}
```

There is only very basic test included