package jsonrpc

import "encoding/json"

type ErrorCode int

const (
	CodeParseError     ErrorCode = -32700
	CodeInvalidRequest           = -32600
	CodeMethodNotFound           = -32601
	CodeInvalidParams            = -32602
	CodeInternalError            = -32603
	CodeMinUserError             = -32099
	CodeMaxUserError             = -32000
)

type Error struct {
	code    ErrorCode
	message string
}

func (error *Error) MarshalJSON() ([]byte, error) {
	marshal := struct {
		Code    ErrorCode `json:"code"`
		Message string    `json:"message"`
	}{Code: error.code, Message: error.message}
	return json.Marshal(marshal)
}

func NewError(code ErrorCode) *Error {
	return &Error{
		code:    code,
		message: code.String(),
	}
}

func NewUserError(code ErrorCode, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

func (code ErrorCode) String() string {
	switch code {
	case CodeParseError:
		return "An error occurred on the server while parsing the JSON text."
	case CodeInvalidRequest:
		return "The JSON sent is not a valid Request object."
	case CodeMethodNotFound:
		return "The method does not exist / is not available."
	case CodeInvalidParams:
		return "Invalid method parameter(s)."
	case CodeInternalError:
		return "Internal JSON-RPC error."
	}
	return "Internal server error."
}
