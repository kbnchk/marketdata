package marketdata

import (
	"reflect"
	"testing"
)

// TODO!

func TestGetData(t *testing.T) {
	source := GarantexNew()
	got, err := source.GetDOM(USDTRUB)
	if err != nil {
		t.Errorf("error getting data = %v", err)
		return
	}
	if reflect.ValueOf(got).IsZero() {
		t.Error("returned empty data")
	}
}
