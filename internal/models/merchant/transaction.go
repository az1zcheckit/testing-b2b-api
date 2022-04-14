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
