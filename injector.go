//go:build wireinject
// +build wireinject

package main

import (
	"go-restful/app"
	"go-restful/controller"
	"go-restful/repository"
	"go-restful/service"
	"net/http"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(repository.NewCategoryRepository, service.NewCategoryService, controller.NewCategoryController)

func InitializedServer() *http.Server {
	wire.Build(app.NewDB, NewValidator, categorySet, app.NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)), NewServer)
	return nil
}
