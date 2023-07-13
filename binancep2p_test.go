package marketdata

import (
	"reflect"
	"testing"
)

func TestBinanceP2P_GetDOM(t *testing.T) {
	source := BinanceP2PNew()
	config := BinanceP2PConfig{
		Fiat:      "TRY",
		Asset:     "USDT",
		Page:      1,
		Rows:      10,
		Countries: []string{"TR"},
		PayTypes:  []string{"VakifBank"},
	}
	got, err := source.GetDOM(config)
	if err != nil {
		t.Errorf("error getting data = %v", err)
		return
	}
	if reflect.ValueOf(got).IsZero() {
		t.Error("returned empty data")
	}

}
