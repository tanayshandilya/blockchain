package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"math/big"
	"math/rand"
	"time"
)

const (
	privateKeyByteSize = 64
)

func RandomBytes(size int) []byte {
	token := make([]byte, size)
	rand.Seed(time.Now().UnixNano())
	rand.Read(token)
	return token
}

func GeneratePrivateKey() []byte {
	return RandomBytes(privateKeyByteSize)
}

func CreateKeyPair() ([]byte, []byte, error) {
	pk := GeneratePrivateKey()
	pkECDSA, err := DecodePrivateKey(pk)
	if err != nil {
		return []byte{}, []byte{}, err
	}
	pubk := EncodePublicKey(&pkECDSA.PublicKey)
	return pk, pubk, nil
}

func DecodePrivateKey(hexBytes []byte) (*ecdsa.PrivateKey, error) {
	hs := hex.EncodeToString(hexBytes)
	pk := new(ecdsa.PrivateKey)
	pk.D, _ = new(big.Int).SetString(hs, 16)
	pk.PublicKey.Curve = elliptic.P256()
	pk.PublicKey.X, pk.PublicKey.Y = pk.PublicKey.Curve.ScalarBaseMult(pk.D.Bytes())
	return pk, nil
}

func EncodePublicKey(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.MarshalCompressed(elliptic.P256(), pub.X, pub.Y)
}

func DecodePublicKey(pub []byte) ecdsa.PublicKey {
	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), pub)
	var key ecdsa.PublicKey
	key.Curve = elliptic.P256()
	key.X = x
	key.Y = y
	return key
}
