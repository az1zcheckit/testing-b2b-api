package merchant

type GetAllAmount struct {
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}

func (a GetAllAmount) Implementation() {}
