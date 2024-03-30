package helper

import (
	"encoding/json"
	"go-restful/model/web"
	"net/http"
)

func ReadRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteResponseBody(writer http.ResponseWriter, response web.Response) {
	writer.WriteHeader(response.Code)
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
