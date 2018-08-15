package main

import (
	"net/url"

	"github.com/satori/go.uuid"
)

type Payment struct {
	Recipient string
	Asset     string
	Amount    string
	Memo      string
}

func (p Payment) PaymentUrl() string {
	u := url.URL{}
	u.Scheme = "mixin"
	u.Path = "pay"

	q := u.Query()
	q.Add("recipient", p.Recipient)
	q.Add("asset", p.Asset)
	q.Add("amount", p.Amount)
	q.Add("memo", p.Memo)
	q.Add("trace", uuid.Must(uuid.NewV4()).String())

	u.RawQuery = q.Encode()
	return u.String()
}
