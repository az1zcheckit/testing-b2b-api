package merchant

type Rates struct {
	Currency string `json:"currency"`
	BuyRate  string `json:"buyRate"`
	SellRate string `json:"sellRate"`
}
