package merchant

import (
	"github.com/godror/godror"
	"time"
)

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

type FilterHistoryTransaction struct {
	Account    []string `json:"account"`
	FromDate   string   `json:"fromDate"`
	ToDate     string   `json:"toDate"`
	Pagination struct {
		CurrentPage int `json:"currentPage"`
		Count       int `json:"count"`
	} `json:"pagination"`
}

type GroupOfTransactions struct {
	GroupDate    string        `json:"groupDate"`
	Transactions []Transaction `json:"transactions"`
}

type Transaction struct {
	Id             string        `json:"id"` // Это DocId
	IdGodror       godror.Number `json:"-"`
	Type           string        `json:"type"`           // TransType
	PaymentPurpose string        `json:"paymentPurpose"` //nazn
	CreatedAtStr   string        `json:"-"`
	CreatedAt      time.Time     `json:"createdAt"` //DateProcess
	ProceededAtStr string        `json:"-"`
	ProceededAt    time.Time     `json:"proceededAt"`
	DocumentNumber string        `json:"documentNumber"`
	Status         string        `json:"status"` // Это DocState

	FromAccount struct {
		Title     string `json:"title,omitempty"`  //senderName
		Number    string `json:"number,omitempty"` //Acc??Number
		ImageName string `json:"imageName,omitempty"`
		Amount    struct {
			Value          string `json:"value,omitempty"` //Amount
			Currency       string `json:"-"`
			CurrencySymbol string `json:"currencySymbol,omitempty"`
		} `json:"amount,omitempty"`
	} `json:"fromAccount,omitempty"`

	ToAccount struct {
		Title  string `json:"title,omitempty"`  // RecipientName
		Number string `json:"number,omitempty"` //Acc??Number
	} `json:"toAccount,omitempty"`
}
