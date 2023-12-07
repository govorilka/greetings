package jsonrpc

import (
	"testing"
)

func TestCodeParseError(t *testing.T) {
	input := NewError(CodeParseError)
	want := "{\"code\":-32700,\"message\":\"An error occurred on the server while parsing the JSON text.\"}"
	checkMarshal(t, input, want)
}

func TestCodeInvalidRequest(t *testing.T) {
	input := NewError(CodeInvalidRequest)
	want := "{\"code\":-32600,\"message\":\"The JSON sent is not a valid Request object.\"}"
	checkMarshal(t, input, want)
}

func TestCodeMethodNotFound(t *testing.T) {
	input := NewError(CodeMethodNotFound)
	want := "{\"code\":-32601,\"message\":\"The method does not exist / is not available.\"}"
	checkMarshal(t, input, want)
}

func TestCodeInvalidParams(t *testing.T) {
	input := NewError(CodeInvalidParams)
	want := "{\"code\":-32602,\"message\":\"Invalid method parameter(s).\"}"
	checkMarshal(t, input, want)
}

func TestCodeInternalError(t *testing.T) {
	input := NewError(CodeInternalError)
	want := "{\"code\":-32603,\"message\":\"Internal JSON-RPC error.\"}"
	checkMarshal(t, input, want)
}

func TestUnknownError(t *testing.T) {
	input := NewError(1)
	want := "{\"code\":1,\"message\":\"Internal server error.\"}"
	checkMarshal(t, input, want)
}

func TestUserError(t *testing.T) {
	input := NewUserError(CodeMinUserError, "Access denied")
	want := "{\"code\":-32099,\"message\":\"Access denied\"}"
	checkMarshal(t, input, want)
}
