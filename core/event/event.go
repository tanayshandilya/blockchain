package event

import (
	"time"

	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
)

type Event struct {
	Version   string      `json:"version"`
	Type      string      `json:"type"`
	Code      string      `json:"code"`
	TimeStamp string      `json:"timeStamp"`
	Data      interface{} `json:"data"`
	Hash      string      `json:"hash"`
}

func (e *Event) create(eventType string, data interface{}) error {
	e.Version = EventVersion
	e.Type = eventType
	e.Code = crypto.HashSHA256([]byte(eventType + EventVersion))
	e.TimeStamp = time.Now().UTC().String()
	e.Data = data
	j, er := encoding.JsonEncode(&e, false)
	if er != nil {
		return er
	}
	e.Hash = crypto.HashSHA256(j)
	return nil
}

func (e *Event) ToJson() ([]byte, error) {
	return encoding.JsonEncode(&e, true)
}

func (e *Event) updateHash() {
	e.Hash = ""
	j, _ := encoding.JsonEncode(&e, false)
	e.Hash = crypto.HashSHA256(j)
}
