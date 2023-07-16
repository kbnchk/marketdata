package marketdata

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_Garantex_GetDOM(t *testing.T) {
	got, err := Garantex().GetDOM("usdtrub")
	if err != nil {
		t.Errorf("error getting data = %v", err)
		return
	}
	if reflect.ValueOf(got).IsZero() {
		t.Error("returned empty data")
	}

}

func TestGarantexHistoryConfig_toParams(t *testing.T) {
	type fields struct {
		Market string
		Limit  int
		From   int
		To     int
		Order  string
	}
	tests := []struct {
		name   string
		fields fields
		want   url.Values
	}{
		{
			name: "market",
			fields: fields{
				Market: "usdtbtc",
			},
			want: url.Values{
				"market": []string{"usdtbtc"},
			},
		},
		{
			name: "market+limit",
			fields: fields{
				Market: "usdtbtc",
				Limit:  1000,
			},
			want: url.Values{
				"market": []string{"usdtbtc"},
				"limit":  []string{"1000"},
			},
		},
		{
			name: "market+limit+from",
			fields: fields{
				Market: "usdtbtc",
				Limit:  1000,
				From:   343434,
			},
			want: url.Values{
				"market": []string{"usdtbtc"},
				"limit":  []string{"1000"},
				"from":   []string{"343434"},
			},
		},
		{
			name: "market+limit+from+to",
			fields: fields{
				Market: "usdtbtc",
				Limit:  1000,
				From:   343434,
				To:     434343,
			},
			want: url.Values{
				"market": []string{"usdtbtc"},
				"limit":  []string{"1000"},
				"from":   []string{"343434"},
				"to":     []string{"434343"},
			},
		},
		{
			name: "market+limit+from+to+order",
			fields: fields{
				Market: "usdtbtc",
				Limit:  1000,
				From:   343434,
				To:     434343,
				Order:  "DESC",
			},
			want: url.Values{
				"market":   []string{"usdtbtc"},
				"limit":    []string{"1000"},
				"from":     []string{"343434"},
				"to":       []string{"434343"},
				"order_by": []string{"DESC"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := GarantexHistoryConfig{
				Market: tt.fields.Market,
				Limit:  tt.fields.Limit,
				From:   tt.fields.From,
				To:     tt.fields.To,
				Order:  tt.fields.Order,
			}
			if got := c.toParams(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GarantexHistoryConfig.toParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
