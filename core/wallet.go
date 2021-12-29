package core

import (
	"errors"
	"fmt"

	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
)

type Wallet struct {
	Address        string `json:"address"`
	PrivateKey     string `json:"privateKey"`
	Balance        int    `json:"balance"`
	StateSignature string `json:"stateSignature"`
}

func (w *Wallet) New() error {
	privateKey, publicKey, err := crypto.CreateKeyPair()
	if err != nil {
		return err
	}
	w.Address = encoding.Base58Encode(publicKey)
	w.PrivateKey = encoding.Base58Encode(privateKey)
	w.Balance = 0
	err1 := w.SignState()
	if err1 != nil {
		return err1
	}
	return nil
}

func (w *Wallet) SignState() error {
	privateKey, err := crypto.DecodePrivateKey(encoding.Base58Decode(w.PrivateKey))
	if err != nil {
		return err
	}
	data := w.Address + fmt.Sprintf("%d", w.Balance)
	sign, err1 := crypto.Sign(privateKey, []byte(data))
	if err1 != nil {
		return err1
	}
	w.StateSignature = encoding.Base58Encode(sign)
	return nil
}

func (w *Wallet) SignTransaction(txn *Transaction) ([]byte, error) {
	privateKey, err := crypto.DecodePrivateKey(encoding.Base58Decode(w.PrivateKey))
	if err != nil {
		return []byte{}, err
	}
	data, err1 := encoding.JsonEncode(txn, false)
	if err1 != nil {
		return []byte{}, err1
	}
	sign, err2 := crypto.Sign(privateKey, data)
	if err2 != nil {
		return []byte{}, err2
	}
	return sign, nil
}

func (w *Wallet) SignBlock(block *Block) ([]byte, error) {
	privateKey, err := crypto.DecodePrivateKey(encoding.Base58Decode(w.PrivateKey))
	if err != nil {
		return []byte{}, err
	}
	data, err1 := encoding.JsonEncode(block, false)
	if err1 != nil {
		return []byte{}, err1
	}
	sign, err2 := crypto.Sign(privateKey, data)
	if err2 != nil {
		return []byte{}, err2
	}
	return sign, nil
}

func (w *Wallet) UpdateBalance(amt int, updateType string) error {
	if updateType == "add" {
		w.Balance = w.Balance + amt
	} else if updateType == "sub" {
		if amt > w.Balance {
			return errors.New("insufficient balance")
		} else {
			w.Balance = w.Balance - amt
		}
	}
	err := w.SignState()
	if err != nil {
		return err
	}
	return nil
}
