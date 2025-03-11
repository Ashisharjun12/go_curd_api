package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/Ashisharjun12/go_curd_api/internal/types"
	"github.com/Ashisharjun12/go_curd_api/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

func New() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err , io.EOF){
			response.WriteJson(w , http.StatusBadRequest , response.GeneralError(err))
			return

		}

		//validator req
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors)
            response.WriteJson(w , http.StatusBadRequest , response.ValidationError(validateErrs))
            return
		}


		slog.Info("creating student api....")
		response.WriteJson(w , http.StatusBadRequest, map[string] string{"success":"Ok"})
	}
}
