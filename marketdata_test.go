package marketdata

import (
	"reflect"
	"testing"
	"time"
)

// TODO!

func TestGetDOM(t *testing.T) {
	source := BinanceNew()
	got, err := source.GetDOM(USDTRUB)
	if err != nil {
		t.Errorf("error getting data = %v", err)
		return
	}
	if reflect.ValueOf(got).IsZero() {
		t.Error("returned empty data")
	}
}

func TestGetHistoryToDate(t *testing.T) {
	source := GarantexNew()
	got, err := source.GetHistoryToDate(USDTRUB, time.Now().Add(-10*time.Minute))
	if err != nil {
		t.Errorf("error getting data = %v", err)
		return
	}
	if reflect.ValueOf(got).IsZero() {
		t.Error("returned empty data")
	}
}

func TestGetHistoryFromID(t *testing.T) {
	source := GarantexNew()
	got, err := source.GetHistoryFromID(USDTRUB, 3840672)
	if err != nil {
		t.Errorf("error getting data = %v", err)
		return
	}
	if reflect.ValueOf(got).IsZero() {
		t.Error("returned empty data")
	}
}
