package main

import (
	"fmt"
	"unsafe"

	"github.com/tanayshandilya/blockchain/core"
	"github.com/tanayshandilya/blockchain/core/crypto"
)

func main() {
	nPrivateKey := crypto.GeneratePrivateKey()

	nPrivateECDSA, pkErr := crypto.DecodePrivateKey(nPrivateKey)

	if pkErr != nil {
		fmt.Println("PrivateKeyErr: ", pkErr)
	}

	nPublicKey := crypto.EncodePublicKey(&nPrivateECDSA.PublicKey)

	nEvent1 := new(core.Event)
	nEvent2 := new(core.Event)

	nEvent1.New("test", "activity_01")
	nEvent2.New("test", "activity_02")

	nTransaction := new(core.Transaction)

	nEventList := new(core.EventList)

	nEventList.Fill(nEvent1, nEvent2)

	txnErr := nTransaction.New("test_txn", 1, nEventList)

	if txnErr != nil {
		fmt.Println("TxnErr: ", txnErr)
	}

	nTxnList := new(core.TransactionList)

	nTxnList.Fill(nTransaction)

	nBlock := new(core.Block)

	bErr := nBlock.New(1, "0000", nTxnList)

	if bErr != nil {
		fmt.Println("BlockErr: ", bErr)
	}

	signErr := nBlock.Sign(nPrivateKey)

	if signErr != nil {
		fmt.Println("SignErr: ", signErr)
	}

	valid, vErr := nBlock.Verify(string(nPublicKey))

	if vErr != nil {
		fmt.Println("VerifyErr: ", vErr)
	}

	if valid {
		fmt.Println("Block Verified")
		j, _ := nBlock.ToJson()
		fmt.Printf("%s", j)
		fmt.Printf("\nsize of block %f m bytes", float32(unsafe.Sizeof(j))/1000000)
	} else {
		fmt.Println("Block Verify Error")
		fmt.Println(nBlock)
	}

}
