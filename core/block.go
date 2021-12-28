package core

import (
	"time"

	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
)

type Block struct {
	Version      string         `json:"version"`
	Height       int            `json:"height"`
	TimeStamp    string         `json:"timestamp"`
	PreviousHash string         `json:"previousHash"`
	Hash         string         `json:"hash"`
	Signature    string         `json:"signature"`
	Transactions []*Transaction `json:"transactions"`
	Consensus    []string       `json:"consensus"`
}

func (b *Block) New(height int, previousHash string, transactions *TransactionList) error {
	b.Version = BlockVersion
	b.Height = height
	b.TimeStamp = time.Now().UTC().Local().String()
	b.PreviousHash = previousHash
	b.Transactions = transactions.Transactions
	j, er := encoding.JsonEncode(b, false)
	if er != nil {
		return er
	}
	b.Hash = crypto.HashSHA256(j)
	return nil
}

func (b *Block) Sign(privateKey []byte) error {
	priK, er := crypto.DecodePrivateKey(privateKey)
	if er != nil {
		return er
	}
	s, er := crypto.Sign(priK, []byte(b.Hash))
	if er != nil {
		return er
	}
	b.Signature = encoding.Base58Encode(s)
	return nil
}

func (b *Block) Verify(address string) (bool, error) {
	pubK := crypto.DecodePublicKey([]byte(address))
	sign := encoding.Base58Decode(b.Signature)
	return crypto.Verify(&pubK, sign, []byte(b.Hash)), nil
}

func (b *Block) ToJson() ([]byte, error) {
	return encoding.JsonEncode(b, true)
}
