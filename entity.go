package marketdata

import (
	"time"
)

type DOM struct {
	MarketPlace string
	MarketName  string
	Date        time.Time
	Bids        []DOMPosition
	Asks        []DOMPosition
}

type DOMPosition struct {
	Price  string
	Amount string
}

type HistoryPosition struct {
	ID     uint
	Date   time.Time
	Price  string // цена
	Volume string // сумма в базовой валюте
	Funds  string // сумма в валюте котировки
}
