package merchant

type Accounts struct {
	Account  	string `json:"account"`
	Balance  	string `json:"balance"`
	Currency    string `json:"currency"`
	CurrencySym string `json:"currencySym"`
	Type	  	string `json:"type"`
	Id	 		string `json:"id"`

}
type AccountDetailsForBankTrans struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
