package jsonrpc

import (
	"bytes"
	"encoding/json"
)

type Params struct {
	named        map[string]json.RawMessage
	positional   []json.RawMessage
	isPositional bool
}

func NewEmptyParams() *Params {
	return &Params{
		named:        make(map[string]json.RawMessage),
		positional:   make([]json.RawMessage, 0),
		isPositional: true,
	}
}

func (params *Params) MarshalJSON() ([]byte, error) {
	if params.isPositional {
		return json.Marshal(params.positional)
	}
	return json.Marshal(params.named)
}

func (params *Params) UnmarshalJSON(data []byte) error {

	var decoder = json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()

	if err := json.Unmarshal(data, &params.named); err == nil {
		params.positional = make([]json.RawMessage, 0)
		params.isPositional = false
		return nil
	}

	if err := json.Unmarshal(data, &params.positional); err == nil {
		params.named = make(map[string]json.RawMessage)
		params.isPositional = true
		return nil
	}

	params.named = make(map[string]json.RawMessage)
	params.positional = make([]json.RawMessage, 0)
	params.isPositional = true
	return NewError(CodeInvalidRequest)
}
