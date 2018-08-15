package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPayment(t *testing.T) {
	p := Payment{
		Recipient: "yiplee",
		Asset:     "btc",
		Amount:    "0.1",
		Memo:      "test",
	}

	u := p.PaymentUrl()
	assert.NotEmpty(t, u)
}
