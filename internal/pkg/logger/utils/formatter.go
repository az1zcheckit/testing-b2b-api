package utils

import "fmt"

type Temp struct {
	method    string
	initiator string
	fn        string
}

func (t Temp) Init(method string, initiator string, fn string) {
	t.method = method
	t.initiator = initiator
	t.fn = fn
}

func (t Temp) ToString() string {
	return fmt.Sprintf("method: %s -> who: %s -> service: %s :::", t.method, t.initiator, t.fn)
}
