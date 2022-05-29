package crypto

import "github.com/tanayshandilya/blockchain/core/encoding"

func Base58Id(size int) string {
	return encoding.Base58Encode(RandomBytes(size))
}
