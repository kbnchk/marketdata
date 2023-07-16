package marketdata

import (
	"time"
)

type DOM struct {
	Date time.Time
	Bids []DOMPosition
	Asks []DOMPosition
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
}
