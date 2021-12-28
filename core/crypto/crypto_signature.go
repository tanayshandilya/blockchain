package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
)

func Sign(privateKey *ecdsa.PrivateKey, data []byte) ([]byte, error) {
	hash := HashSHA256(data)
	sign, serr := ecdsa.SignASN1(rand.Reader, privateKey, []byte(hash))
	if serr != nil {
		return []byte{}, serr
	}

	return sign, nil
}

func Verify(publicKey *ecdsa.PublicKey, signature []byte, data []byte) bool {
	hash := HashSHA256(data)
	return ecdsa.VerifyASN1(publicKey, []byte(hash), signature)
}
