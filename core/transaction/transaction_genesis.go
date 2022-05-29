package transaction

import "github.com/tanayshandilya/blockchain/core/event"

func GenesisTransaction() *Transaction {
	e := event.GenesisEvent()
	t := new(Transaction)
	t.New(txn_GenesisType, []*event.Event{e})
	t.TimeStamp = txn_GenesisTimestamp
	t.updateHash()
	return t
}
