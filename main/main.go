package main

import (
	"fmt"
	"io"
	"net/http"

	"example/dictionary"
)

func value(w http.ResponseWriter, req *http.Request, values *dictionary.Dictionary) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Empty name", http.StatusBadRequest)
		return
	}
	result, err := values.Get(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, result)
}

func main() {
	values := dictionary.New()
	http.HandleFunc("/value", func(w http.ResponseWriter, req *http.Request) {
		value(w, req, values)
	})
	http.ListenAndServe(":8090", nil)
}
