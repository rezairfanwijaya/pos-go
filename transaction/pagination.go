package transaction

type PaginationTransaction struct {
	Limit        int         `json:"limit"`
	Page         int         `json:"page"`
	TotalData    int         `json:"total_data"`
	TotalPage    int         `json:"total_page"`
	Transactions interface{} `json:"transactions"`
	FirstPage    string      `json:"first_page"`
	PreviousPage string      `json:"previous_page"`
	NextPage     string      `json:"next_page"`
	LastPage     string      `json:"last_page"`
}
