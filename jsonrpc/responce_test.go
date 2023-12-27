package jsonrpc

import "testing"

func TestResponceMethodNotFoundError(t *testing.T) {
	input := NewResponceByError(NewIntId(1), NewError(CodeMethodNotFound))
	want := "{\"jsonrpc\":\"2.0\",\"error\":{\"code\":-32601,\"message\":\"The method does not exist / is not available.\"},\"id\":1}"
	checkMarshal(t, input, want)
}

func TestResponceParseErrorError(t *testing.T) {
	input := NewResponceByError(nil, NewError(CodeParseError))
	want := "{\"jsonrpc\":\"2.0\",\"error\":{\"code\":-32700,\"message\":\"Parse error\"},\"id\":null}"
	checkMarshal(t, input, want)
}

func TestResponceValue(t *testing.T) {
	input := NewResponceByValue(NewIntId(1), "value")
	want := "{\"jsonrpc\":\"2.0\",\"result\":\"value\",\"id\":1}"
	checkMarshal(t, input, want)
}

func TestResponceInvalidID(t *testing.T) {
	input := NewResponceByValue(nil, "value")
	want := "{\"jsonrpc\":\"2.0\",\"error\":{\"code\":-32603,\"message\":\"Internal JSON-RPC error.\"},\"id\":null}"
	checkMarshal(t, input, want)
}
