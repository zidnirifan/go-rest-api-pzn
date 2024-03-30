package exception

import (
	"fmt"
	"go-restful/helper"
	"go-restful/model/web"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type RequestError struct {
	StatusCode int
	Err        error
}

func (r *RequestError) Error() string {
	return r.Err.Error()
}

func HandleErrorRequest(err error, writer http.ResponseWriter) {
	exception, ok := err.(*RequestError)
	if ok {
		response := web.Response{
			Ok:      false,
			Code:    exception.StatusCode,
			Message: exception.Error(),
		}
		helper.WriteResponseBody(writer, response)
		return
	}

	validationErrors, ok := err.(validator.ValidationErrors)
	if ok {
		for _, e := range validationErrors {
			errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			response := web.Response{
				Ok:      false,
				Code:    http.StatusBadRequest,
				Message: errorMessage,
			}
			helper.WriteResponseBody(writer, response)
			return
		}
	}

	response := web.Response{
		Ok:      false,
		Code:    http.StatusInternalServerError,
		Message: exception.Error(),
	}
	helper.WriteResponseBody(writer, response)
}
