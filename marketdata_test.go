package marketdata

import (
	"reflect"
	"testing"
)

// TODO!

func TestGetData(t *testing.T) {
	source, err := NewSource(Beribit)
	if err != nil {
		t.Errorf("MarketDataSource creating error = %v", err)
		return
	}
	got, err := source.GetDOM(USDTRUB)
	if err != nil {
		t.Errorf("GetGarantexData() error = %v", err)
		return
	}
	if reflect.DeepEqual(got, DOM{}) {
		t.Error("GetGarantexData() returned empty data")
	}
}
