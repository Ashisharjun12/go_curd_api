package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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


func ValidationError(errs validator.ValidationErrors) Response {
	var errMsgs []string


	for _,err := range errs{
		switch err.ActualTag(){
			case "required":
                errMsgs = append(errMsgs, err.Field()+" is required")
            default:
                errMsgs = append(errMsgs, err.Field()+" has an invalid value")
		}

	}
	return Response{Status: "error", Error: strings.Join(errMsgs, ", ")}
}
