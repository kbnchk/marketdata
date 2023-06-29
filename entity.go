package marketdata

import "time"

type DOM struct {
	MarketPlace string
	MarketName  string
	Date        time.Time
	Bids        []DOMPosition
	Asks        []DOMPosition
}

type DOMPosition struct {
	Price  float64
	Amount float64
	Type   string
	Factor float64
}

type HistoryPosition struct {
	ID     uint
	Date   time.Time
	Price  float32 // цена
	Volume float32 // сумма в базовой валюте
	Funds  float32 // сумма в валюте котировки
}
