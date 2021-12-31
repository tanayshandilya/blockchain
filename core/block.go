package core

import (
	"strings"
	"time"

	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
)

type Block struct {
	Version      string         `json:"version"`
	Height       int            `json:"height"`
	PreviousHash string         `json:"previousHash"`
	TimeStamp    string         `json:"timestamp"`
	Hash         string         `json:"hash"`
	Markle       string         `json:"markle"`
	Signature    string         `json:"signature"`
	Transactions []*Transaction `json:"transactions"`
	Consensus    []string       `json:"consensus"`
}

func (b *Block) New(height int, previousHash string, transactions []*Transaction) error {
	b.Version = BlockVersion
	b.Height = height
	b.TimeStamp = time.Now().UTC().String()
	b.PreviousHash = previousHash
	b.Transactions = transactions
	b.Markle = createTxnMarkle(transactions)
	j, er := encoding.JsonEncode(b, false)
	if er != nil {
		return er
	}
	b.Hash = crypto.HashSHA256(j)
	return nil
}

func createTxnMarkle(txns []*Transaction) string {
	hashes := []string{}
	for _, t := range txns {
		hashes = append(hashes, t.Hash)
	}
	return crypto.HashSHA512([]byte(strings.Join(hashes, ".")))
}

func (b *Block) ToJson() ([]byte, error) {
	return encoding.JsonEncode(b, true)
}

func (b *Block) createGenesis() {
	e := new(Event)
	t := new(Transaction)
	e.New("genesis", "origin")
	e.TimeStamp = "1996-05-31 00:00:00.0000000 +0000 UTC"
	e.updateHash()
	t.New("genesis", 0, []*Event{e})
	t.TimeStamp = "1996-05-31 00:00:00.0000000 +0000 UTC"
	t.updateHash()
	b.Version = BlockVersion
	b.Height = 0
	b.TimeStamp = "1996-05-31 00:00:00.0000000 +0000 UTC"
	b.PreviousHash = "0"
	b.Transactions = []*Transaction{t}
	b.Markle = createTxnMarkle([]*Transaction{t})
	j, _ := encoding.JsonEncode(b, false)
	b.Hash = crypto.HashSHA256(j)
}
