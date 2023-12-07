package jsonrpc

import (
	"testing"
)

func TestStringId(t *testing.T) {
	input := NewStringId("100")
	want := "\"100\""
	checkMarshal(t, input, want)
}

func TestIntId(t *testing.T) {
	input := NewIntId(100)
	want := "100"
	checkMarshal(t, input, want)
}
