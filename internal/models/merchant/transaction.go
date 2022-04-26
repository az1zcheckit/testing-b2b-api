package merchant

type Transactions struct {
	AccountKt     string `json:"accKtNumber"`
	AccountDt     string `json:"accDtNumber"`
	DocId         string `json:"docId"`
	DateProcess   string `json:"dateProcess"`
	Nazn          string `json:"nazn"`
	TransType     string `json:"transType"`
	SenderName    string `json:"senderName"`
	RecipientName string `json:"recipientName"`
	DocState      string `json:"docState"`
	Amount        string `json:"summa"`
}

type Transaction struct {
	CreatedDate    string `json:"createdDate"`
	ProceedDate    string `json:"proceedDate"`
	AccFrom        string `json:"accFrom"`
	AccTo          string `json:"accTo"`
	TransId        string `json:"transId"`
	TransType      string `json:"transType"`
	RecipientName  string `json:"recipientName"`
	PaymentPurpose string `json:"paymentPurpose"`
	Amount         string `json:"amount"`
}
