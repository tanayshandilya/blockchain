package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"

	"github.com/tanayshandilya/blockchain/core/encoding"
)

func HashSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func HashSHA512(data []byte) string {
	hash := sha512.Sum512(data)
	return encoding.Base58Encode(hash[:])
}
