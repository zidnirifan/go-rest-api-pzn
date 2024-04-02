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
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(response.Code)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
