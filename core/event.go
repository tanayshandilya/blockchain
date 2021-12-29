package core

import (
	"time"

	"github.com/tanayshandilya/blockchain/core/crypto"
	"github.com/tanayshandilya/blockchain/core/encoding"
)

type Event struct {
	Type      string `json:"type"`
	TimeStamp string `json:"timeStamp"`
	Data      string `json:"data"`
	Hash      string `json:"hash"`
}

type EventList struct {
	Events []*Event `json:"events"`
}

func (e *Event) New(eventType string, data string) error {
	e.Type = eventType
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

func (e *EventList) Fill(events ...*Event) {
	e.Events = events
}
