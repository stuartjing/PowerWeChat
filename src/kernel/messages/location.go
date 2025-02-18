package messages

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
)

type Location struct {
	*Message
}

func NewLocation(items *power.HashMap) *Location {
	m := &Location{
		NewMessage(items),
	}
	m.Type = "location"

	m.Properties = []string{
		"latitude",
		"longitude",
		"scale",
		"label",
		"precision",
	}

	return m
}
