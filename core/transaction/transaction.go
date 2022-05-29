package transaction

import (
	"strings"
	"time"

	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
	"github.com/tanayshandilya/blockchain/core/event"
)

type Transaction struct {
	Version   string         `json:"version"`
	Type      string         `json:"type"`
	TimeStamp string         `json:"timeStamp"`
	Hash      string         `json:"hash"`
	Markle    string         `json:"markle"`
	Events    []*event.Event `json:"events"`
}

func (t *Transaction) New(txnType string, events []*event.Event) error {
	t.Version = TransactionVersion
	t.Type = txnType
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

func createEventMarkle(events []*event.Event) string {
	hashes := []string{}
	for _, e := range events {
		hashes = append(hashes, e.Hash)
	}
	return crypto.HashSHA512([]byte(strings.Join(hashes, ".")))
}

func (t *Transaction) ToJson() ([]byte, error) {
	return encoding.JsonEncode(&t, true)
}

func (t *Transaction) updateHash() {
	t.Hash = ""
	j, _ := encoding.JsonEncode(&t, false)
	t.Hash = crypto.HashSHA256(j)
}
