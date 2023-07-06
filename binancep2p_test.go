package marketdata

import (
	"reflect"
	"testing"
)

func TestBinanceP2P_GetDOM(t *testing.T) {
	source := BinanceP2PNew()
	got, err := source.GetDOM(USDTTHB)
	if err != nil {
		t.Errorf("error getting data = %v", err)
		return
	}
	if reflect.ValueOf(got).IsZero() {
		t.Error("returned empty data")
	}

}
