package merchant

type ConversionRequest struct {
	AccountFrom string  `json:"accFrom"`
	AccountTo   string  `json:"accTo"`
	Amount  	float64 `json:"amount"`
	Footnote    string  `json:"footnote"`
	Rate 		float32 `json:"rate"`
	DocNumber   string	`json:"docNumber"`
	ExecutionDate string `json:"executionDate"`
}
type Amount struct{
	Value 		   float64 `json:"value"`
	CurrencySymbol string  `json:"currencySymbol"`
}
type Account struct{
	Name string `json:"name"`
	Currency string `json:"currency"`
}

type Signature struct{
	FullName string `json:"fullName"`
	Role string	`json:"role"`
	Status string `json:"status"`
}

type ConversionResponse struct{
	Id				 string  `json:"id"`
	DocumentNumber   string	 `json:"documentNumber"`
	Rate 		     string  `json:"rate"`
	Footnote      	 string  `json:"footnote"`
	ExecutionDate 	 string  `json:"executionDate"`
	CreateDate 	     string  `json:"createDate"`
	AmountFrom    	 Amount  `json:"amountFrom"`
	AmountTo      	 Amount  `json:"amountTo"`
	AccountFrom   	 Account `json:"accountFrom"`
	AccountTo     	 Account `json:"accountTo"`
	Signature	 []Signature `json:"signature"`
}

type ConvertedValue struct{
	Value 		 float64 `json:"value"`
	Rate 		 string  `json:"rate"`
}

func GetSymCurrency(currency string)string{
	switch currency{
	case "TJS":
		return "c"
	case "USD":
		return "$"
	case "RUB":
		return "₽"
	case "EUR":
		return "€"
	default:
		return "currency not found"
	}
}

