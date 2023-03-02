package transaction

type InputTransaction struct {
	Amount int64  `json:"amount" binding:"required"`
	Notes  string `json:"notes" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Type   string `json:"type" binding:"required"`
}
