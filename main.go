package main

import (
	"go-restful/helper"
	"go-restful/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func NewServer(router http.Handler) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
}

func NewValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func main() {
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
