package main

import (
	"io"
	"net/http"

	"example/dictionary"
	"example/jsonrpc"
)

func value(w http.ResponseWriter, req *http.Request, values *dictionary.Dictionary) {
	id := jsonrpc.NewStringId("1")

	body, err := io.ReadAll(req.Body)
	if err != nil {
		jsonrpc.WriteError(w, id, jsonrpc.CodeInvalidParams, "Empty name")
		return
	}

	result, err := values.Get(string(body))
	if err != nil {
		jsonrpc.WriteError(w, id, jsonrpc.CodeInvalidRequest, err.Error())
		return
	}
	jsonrpc.WriteResult(w, id, result)
}

func main() {
	values := dictionary.New()
	http.HandleFunc("/value", func(w http.ResponseWriter, req *http.Request) {
		value(w, req, values)
	})
	http.ListenAndServe(":8090", nil)
}
