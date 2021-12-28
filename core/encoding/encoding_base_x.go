package encoding

import (
	"encoding/base64"

	"github.com/btcsuite/btcutil/base58"
)

func Base58Encode(data []byte) string {
	return base58.Encode(data)
}

func Base58Decode(encoded string) []byte {
	return base58.Decode(encoded)
}

func Base64Encode(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func Base64Decode(encoded string) ([]byte, error) {
	data, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}
