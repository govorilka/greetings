package jsonrpc

import (
	"bytes"
	"encoding/json"
)

type Request struct {
	id     *Id
	method string
	params *Params
}

func NewEmptyRequest() *Request {
	return &Request{
		id: nil,
	}
}

func NewRequest(id *Id) *Request {
	return &Request{
		id: id,
	}
}

func (request *Request) MarshalJSON() ([]byte, error) {
	marshal := struct {
		Jsonrpc string  `json:"jsonrpc"`
		Method  string  `json:"method"`
		Params  *Params `json:"params"`
		Id      *Id     `json:"id,omitempty"`
	}{Jsonrpc: "2.0", Id: request.id, Method: request.method, Params: request.params}
	return json.Marshal(marshal)
}

func (request *Request) UnmarshalJSON(data []byte) error {
	var result struct {
		Jsonrpc string  `json:"jsonrpc"`
		Method  string  `json:"method"`
		Params  *Params `json:"params"`
		Id      *Id     `json:"id"`
	}

	var decoder = json.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&result); err != nil {
		return err
	}

	request.id = result.Id
	request.method = result.Method
	request.params = result.Params
	return nil
}

func (request *Request) Decode(data []byte) *Error {
	if err := json.NewDecoder(bytes.NewReader(data)).Decode(&request); err != nil {
		if rpcError, ok := err.(*Error); ok {
			return rpcError
		}
		return NewError(CodeParseError)
	}
	return nil
}
