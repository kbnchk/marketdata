package marketdata

import (
	"reflect"
	"testing"
	"time"
)

func Test_garantexHistoryPosition_toEntity(t *testing.T) {
	tests := []struct {
		name     string
		position garantexHistoryPosition
		want     HistoryPosition
	}{
		{
			name: "default",
			position: garantexHistoryPosition{
				ID:     float64(3834759),
				Date:   "2023-06-29T13:30:07+03:00",
				Price:  "88.48",
				Volume: "10023.05",
				Funds:  "886839.46",
			},
			want: HistoryPosition{
				ID:     uint(3834759),
				Date:   time.Date(2023, 6, 29, 13, 30, 07, 0, time.Local),
				Price:  float32(88.48),
				Volume: float32(10023.05),
				Funds:  float32(886839.46),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.position.toEntity(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("garantexHistoryPosition.toEntity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_garantex_GetHistoryToDate(t *testing.T) {
	garantex := Garantex{
		domURL:     "https://garantex.io/api/v2/depth",
		historyURL: "https://garantex.io/api/v2/trades",
	}

	earliest := time.Now().Add(-24 * time.Hour)
	got, err := garantex.GetHistoryToDate(USDTRUB, earliest)
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
