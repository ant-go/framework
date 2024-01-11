package stdjson

import (
	"encoding/json"
)

type StdJSON struct{}

func (StdJSON) Marshal(value any) (data []byte, err error) {
	return json.Marshal(value)
}

func (StdJSON) Unmarshal(data []byte, value any) (err error) {
	return json.Unmarshal(data, value)
}
