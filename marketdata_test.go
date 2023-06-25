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
		t.Errorf("error getting data = %v", err)
		return
	}
	if reflect.DeepEqual(got, DOM{}) {
		t.Error("returned empty data")
	}
}
