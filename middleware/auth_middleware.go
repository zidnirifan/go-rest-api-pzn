package middleware

import (
	"go-restful/helper"
	"go-restful/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "RAHASIA" {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		response := web.Response{
			Ok:      false,
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		}
		helper.WriteResponseBody(writer, response)
	}
}
