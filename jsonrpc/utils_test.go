package jsonrpc

import (
	"encoding/json"
	"testing"
)

func checkMarshal(t *testing.T, input json.Marshaler, want string) {
	got, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	result := string(got)
	if result != want {
		t.Errorf("got %q, wanted %q", result, want)
	}
}
