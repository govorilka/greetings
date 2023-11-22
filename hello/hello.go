package main

import (
	"fmt"
	"io"
	"net/http"

	"example.com/greetings"
)

func hello(w http.ResponseWriter, req *http.Request, values *greetings.Values) {
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
	values := greetings.New()
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		hello(w, req, values)
	})
	http.ListenAndServe(":8090", nil)
}
