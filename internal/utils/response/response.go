package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string
	Error  string
}

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) Response {
	return Response{Status: "error", Error: err.Error()}

}
