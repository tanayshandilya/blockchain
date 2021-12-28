package encoding

import (
	"encoding/hex"
	"errors"
)

func HexEncode(data []byte) string {
	return hex.EncodeToString(data)
}

func Hex0xEncode(data []byte) string {
	return "0x" + hex.EncodeToString(data)
}

func HexDecode(encoded string) ([]byte, error) {
	data, err := hex.DecodeString(encoded)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func Hex0xDecode(encoded string) ([]byte, error) {
	zeroX := encoded[:2]
	if zeroX != "0x" {
		return []byte{}, errors.New("invalid 0x prefix")
	}
	data, err := hex.DecodeString(encoded[2:])
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}
