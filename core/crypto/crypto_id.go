package crypto

import "github.com/tanayshandilya/blockchain/core/encoding"

func EventId() string {
	return encoding.Base58Encode(RandomBytes(8))
}

func TransactionId() string {
	return encoding.Base58Encode(RandomBytes(12))
}
