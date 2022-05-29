package asset

import (
	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
)

type Question struct {
	Id             string  `json:"id"`
	Timestamp      string  `json:"timestamp"`
	Status         string  `json:"status"`
	Owner          string  `json:"owner"`
	Royalty        float64 `json:"royalty"`
	IsTransferable bool    `json:"isTransferable"`
	Markle         string  `json:"markle"`
	Hash           string  `json:"hash"`
}

type License struct {
	Id          string      `json:"id"`
	Timestamp   string      `json:"timestamp"`
	Issuer      string      `json:"issuer"`
	IssuedTo    string      `json:"issuedTo"`
	Expires     string      `json:"expires"`
	Assets      string      `json:"assets"`
	Markle      string      `json:"markle"`
	Certificate Certificate `json:"certificate"`
	Hash        string      `json:"hash"`
}

type Certificate struct {
	Id                     string `json:"id"`
	Version                string `json:"version"`
	SerialNumber           string `json:"serialNumber"`
	SignatureAlgorithm     string `json:"signatureAlgorithm"`
	SignatureHashAlgorithm string `json:"signatureHashAlgorithm"`
	Issuer                 string `json:"issuer"`
	ValidFrom              string `json:"validFrom"`
	ValidTill              string `json:"validTill"`
	Subject                string `json:"subject"`
	PublicKey              string `json:"publicKey"`
	PublicKeyAlgorithm     string `json:"publicKeyAlgorithm"`
}

func (q *Question) UpdateHash() {
	q.Hash = ""
	j, _ := encoding.JsonEncode(&q, false)
	q.Hash = crypto.HashSHA256(j)
}
