package jsonrpc

import (
	"bytes"
	"encoding/json"
	"testing"
)

func checkRequest(t *testing.T, input []byte, want string) {
	request := NewEmptyRequest()
	if err := json.NewDecoder(bytes.NewReader(input)).Decode(&request); err != nil {
		t.Fatal(err)
		return
	}

	got, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}

	result := string(got)
	if result != want {
		t.Fatalf("got %q, wanted %q", result, want)
	}
}

func checkRequestError(t *testing.T, input []byte, want string) {
	request := NewEmptyRequest()
	if err := request.Decode(input); err != nil {
		if err.Error() != want {
			t.Fatalf("got %q, wanted %q", err.Error(), want)
		}
		return
	}
	t.Fatalf("No request error")
}

func TestRequestCall1(t *testing.T) {
	input := []byte(`{"jsonrpc": "2.0", "method": "subtract", "params": [42, 23], "id": 1}`)
	want := "{\"jsonrpc\":\"2.0\",\"method\":\"subtract\",\"params\":[42,23],\"id\":1}"
	checkRequest(t, input, want)
}

func TestRequestCall2(t *testing.T) {
	input := []byte(`{"jsonrpc": "2.0", "method": "subtract", "params": {"subtrahend": 23, "minuend": 42}, "id": 3}`)
	want := "{\"jsonrpc\":\"2.0\",\"method\":\"subtract\",\"params\":{\"minuend\":42,\"subtrahend\":23},\"id\":3}"
	checkRequest(t, input, want)
}

func TestRequestNotification(t *testing.T) {
	input := []byte(`{"jsonrpc": "2.0", "method": "update", "params": [1,2,3,4,5]}`)
	want := "{\"jsonrpc\":\"2.0\",\"method\":\"update\",\"params\":[1,2,3,4,5]}"
	checkRequest(t, input, want)
}

func TestRequestParseError(t *testing.T) {
	input := []byte(`{"jsonrpc": "2.0", "method": "foobar, "params": "bar", "baz]`)
	want := "Parse error"
	checkRequestError(t, input, want)
}

func TestRequestParseError2(t *testing.T) {
	input := []byte(`[]`)
	want := "Parse error"
	checkRequestError(t, input, want)
}

func TestRequestParseError3(t *testing.T) {
	input := []byte(`[1]`)
	want := "Parse error"
	checkRequestError(t, input, want)
}

func TestRequestInvalidRequest(t *testing.T) {
	input := []byte(`{"jsonrpc": "2.0", "method": 1, "params": "bar"}`)
	want := "Invalid request"
	checkRequestError(t, input, want)
}
