package jsonrpc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteResult(w http.ResponseWriter, id *Id, result string) {
	got, err := json.Marshal(NewResponceByValue(id, result))
	if err != nil {
		got, err := json.Marshal(NewInternalErrorResponce(id, err.Error()))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			http.Error(w, string(got), http.StatusInternalServerError)
		}
		return
	}
	fmt.Fprintf(w, string(got))
}

func WriteError(w http.ResponseWriter, id *Id, code ErrorCode, error string) {
	got, err := json.Marshal(NewResponceByError(id, NewUserError(code, error)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, string(got), http.StatusInternalServerError)
}
