package merchant

type Conversion struct {
	AccFrom string  `json:"accFrom"`
	AccTo   string  `json:"accTo"`
	Amount  float64 `json:"amount"`
	Dest    string  `json:"dest"`
}
