package main

import (
	"fmt"
	"go-restful/config"
	"go-restful/helper"
	"go-restful/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func NewServer(router http.Handler) *http.Server {
	port := config.GetConfig().Port

	return &http.Server{
		Addr:    fmt.Sprintf("localhost:%s", port),
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
