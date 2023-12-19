package jsonrpc

import "encoding/json"

type Responce struct {
	id     *Id
	error  *Error
	result string
}

func NewResponceByValue(id *Id, result string) *Responce {
	if id == nil {
		return &Responce{
			id:     id,
			error:  NewError(CodeInternalError),
			result: "",
		}
	}
	return &Responce{
		id:     id,
		error:  nil,
		result: result,
	}
}

func NewResponceByError(id *Id, error *Error) *Responce {
	return &Responce{
		id:     id,
		error:  error,
		result: "",
	}
}

func NewInternalErrorResponce(id *Id, error string) *Responce {
	return &Responce{
		id:     id,
		error:  NewUserError(CodeInternalError, error),
		result: "",
	}
}

func (responce *Responce) MarshalJSON() ([]byte, error) {
	if responce.error != nil {
		marshal := struct {
			Jsonrpc string `json:"jsonrpc"`
			Error   *Error `json:"error"`
			ID      *Id    `json:"id"`
		}{Jsonrpc: "2.0", Error: responce.error, ID: responce.id}
		return json.Marshal(marshal)
	}
	marshal := struct {
		Jsonrpc string `json:"jsonrpc"`
		Result  string `json:"result"`
		ID      *Id    `json:"id"`
	}{Jsonrpc: "2.0", Result: responce.result, ID: responce.id}
	return json.Marshal(marshal)
}
