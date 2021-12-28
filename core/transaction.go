package core

import (
	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
)

type Transaction struct {
	Id     string   `json:"id"`
	Type   string   `json:"type"`
	Code   int      `json:"code"`
	Hash   string   `json:"hash"`
	Events []*Event `json:"events"`
}

type TransactionList struct {
	Transactions []*Transaction `json:"transactions"`
}

func (t *Transaction) New(txnType string, txnCode int, events *EventList) error {
	t.Id = crypto.TransactionId()
	t.Type = txnType
	t.Code = txnCode
	t.Events = events.Events
	j, er := encoding.JsonEncode(&t, false)
	if er != nil {
		return er
	}
	t.Hash = crypto.HashSHA256(j)
	return nil
}

func (t *Transaction) ToJson() ([]byte, error) {
	return encoding.JsonEncode(&t, true)
}

func (t *TransactionList) Fill(txns ...*Transaction) {
	t.Transactions = txns
}
