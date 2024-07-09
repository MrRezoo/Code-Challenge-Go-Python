package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
)

func MethodHandler(method string, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			err := errors.New("method not allowed")
			response := GenerateBaseResponseWithError(nil, false, http.StatusMethodNotAllowed, err)
			w.WriteHeader(response.Code)
			json.NewEncoder(w).Encode(response)
			return
		}
		handlerFunc(w, r)
	}
}
