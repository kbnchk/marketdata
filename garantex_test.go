package marketdata

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_Garantex_GetDOM(t *testing.T) {
	got, err := GarantexNew().GetDOM("usdtrub")
	if err != nil {
		t.Errorf("error getting data = %v", err)
		return
	}
	if reflect.ValueOf(got).IsZero() {
		t.Error("returned empty data")
	}

}

func Test_garantexGetHistory(t *testing.T) {
	var params url.Values
	params.Add("market", "usdtbtc")
	params.Add("limit", "1000")
	got, err := GarantexNew().GetHistory(params)
	if err != nil {
		t.Errorf("garantex.GetHistoryToDate() error = %v", err)
		return
	}
	if len(got) == 0 {
		t.Error("garantex.GetHistoryToDate() returned empty data")
		return
	}
	for i := 0; i < len(got)-2; i++ {
		if got[i].ID <= got[i+1].ID {
			t.Errorf("garantex.GetHistoryToDate() have data order errors: el[%d].ID=%d, and el[%d].ID=%d (descending order)", i, got[i].ID, i+1, got[i+1].ID)
		}
	}
}
