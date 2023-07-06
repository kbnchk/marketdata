package marketdata

import (
	"reflect"
	"testing"
	"time"
)

func Test_binanceDOMResponse_toEntity(t *testing.T) {
	tests := []struct {
		name       string
		response   binanceDOMResponse
		markettype MarketType
		want       DOM
	}{
		{
			name: "success",
			response: binanceDOMResponse{
				Bids: [][]string{
					{"93.18000000", "570.00000000"},
					{"93.16000000", "56223.00000000"},
				},
				Asks: [][]string{
					{"93.19000000", "971.00000000"},
					{"93.20000000", "3975.00000000"},
				},
			},
			markettype: USDTRUB,
			want: DOM{
				Date:        time.Now().UTC().Round(1 * time.Second),
				MarketPlace: "binance",
				MarketName:  "USDT/RUB",
				Bids: []DOMPosition{
					{
						Price:  "93.18000000",
						Amount: "570.00000000",
					},
					{
						Price:  "93.16000000",
						Amount: "56223.00000000",
					},
				},
				Asks: []DOMPosition{
					{
						Price:  "93.19000000",
						Amount: "971.00000000",
					},
					{
						Price:  "93.20000000",
						Amount: "3975.00000000",
					},
				},
			},
		},

		//TODO
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.response.toEntity(tt.markettype)
			got.Date = got.Date.Round(1 * time.Second)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binanceDOMResponse.toEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}
