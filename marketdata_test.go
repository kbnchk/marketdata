package marketdata

import (
	"reflect"
	"testing"
)

// TODO!

func TestGetGarantexData(t *testing.T) {
	got, err := GetGarantexData(USDTRUB)
	if err != nil {
		t.Errorf("GetGarantexData() error = %v", err)
		return
	}
	if reflect.DeepEqual(got, Response{}) {
		t.Error("GetGarantexData() returned empty data")
	}
}

func TestGetBeribitData(t *testing.T) {
	got, err := GetBeribitData(USDTRUB)
	if err != nil {
		t.Errorf("GetGarantexData() error = %v", err)
		return
	}
	if reflect.DeepEqual(got, Response{}) {
		t.Error("GetGarantexData() returned empty data")
	}
}
