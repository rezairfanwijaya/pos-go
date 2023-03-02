package transaction

import "time"

type Transaction struct {
	ID     int       `json:"id" gorm:"primaryKey"`
	Amount int64     `json:"amount"`
	Notes  string    `json:"notes"`
	Date   time.Time `json:"date"`
	Type   string    `json:"type"`
}
