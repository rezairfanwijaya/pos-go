package transaction

type InputNewTransaction struct {
	Amount int64  `json:"amount" binding:"required"`
	Notes  string `json:"notes" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Type   string `json:"type" binding:"required"`
}

type InputEditTransaction struct {
	Amount int64  `json:"amount" binding:"required"`
	Notes  string `json:"notes" binding:"required"`
	Date   string `json:"date" binding:"required"`
	Type   string `json:"type" binding:"required"`
}

type ParamsGetAllTransaction struct {
	Amount          string
	Date            string
	TransactionType string
	FromAmount      int
	ToAmount        int
	Limit           int
	Page            int
}
