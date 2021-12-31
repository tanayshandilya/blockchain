package core

import (
	"strings"
	"time"

	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
)

type Transaction struct {
	Version   string   `json:"version"`
	Type      string   `json:"type"`
	Code      int      `json:"code"`
	TimeStamp string   `json:"timeStamp"`
	Hash      string   `json:"hash"`
	Markle    string   `json:"markle"`
	Events    []*Event `json:"events"`
}

type TransactionList struct {
	Transactions []*Transaction `json:"transactions"`
}

func (t *Transaction) New(txnType string, txnCode int, events []*Event) error {
	t.Version = TransactionVersion
	t.Type = txnType
	t.Code = txnCode
	t.Events = events
	t.TimeStamp = time.Now().UTC().String()
	t.Markle = createEventMarkle(events)
	j, er := encoding.JsonEncode(&t, false)
	if er != nil {
		return er
	}
	t.Hash = crypto.HashSHA256(j)
	return nil
}

func createEventMarkle(events []*Event) string {
	hashes := []string{}
	for _, e := range events {
		hashes = append(hashes, e.Hash)
	}
	return crypto.HashSHA512([]byte(strings.Join(hashes, ".")))
}

func (t *Transaction) ToJson() ([]byte, error) {
	return encoding.JsonEncode(&t, true)
}

func (t *TransactionList) Fill(txns ...*Transaction) {
	t.Transactions = txns
}

func (t *Transaction) updateHash() {
	t.Hash = ""
	j, _ := encoding.JsonEncode(&t, false)
	t.Hash = crypto.HashSHA256(j)
}
