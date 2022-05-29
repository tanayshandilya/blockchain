package block

import (
	"strings"
	"time"

	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
	"github.com/tanayshandilya/blockchain/core/transaction"
)

type Block struct {
	Version      string                     `json:"version"`
	Height       int                        `json:"height"`
	PreviousHash string                     `json:"previousHash"`
	TimeStamp    string                     `json:"timestamp"`
	Hash         string                     `json:"hash"`
	Markle       string                     `json:"markle"`
	Signature    string                     `json:"signature"`
	Consensus    []string                   `json:"consensus"`
	Transactions []*transaction.Transaction `json:"transactions"`
}

func (b *Block) New(height int, previousHash string, transactions []*transaction.Transaction) error {
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

func createTxnMarkle(txns []*transaction.Transaction) string {
	hashes := []string{}
	for _, t := range txns {
		hashes = append(hashes, t.Hash)
	}
	return crypto.HashSHA512([]byte(strings.Join(hashes, ".")))
}

func (b *Block) ToJson() ([]byte, error) {
	return encoding.JsonEncode(b, true)
}

func (b *Block) AddConsensus() {

}

func (b *Block) CreateGenesis() {
	t := transaction.GenesisTransaction()
	b.Version = BlockVersion
	b.Height = 0
	b.TimeStamp = block_GenesisTimestamp
	b.PreviousHash = block_GenesisPrevHash
	b.Transactions = []*transaction.Transaction{t}
	b.Markle = createTxnMarkle([]*transaction.Transaction{t})
	j, _ := encoding.JsonEncode(b, false)
	b.Hash = crypto.HashSHA256(j)
}
