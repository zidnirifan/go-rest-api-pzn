package app

import (
	"go-restful/controller"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/categories", controller.FindAll)
	router.GET("/api/categories/:categoryId", controller.FindById)
	router.POST("/api/categories", controller.Create)
	router.PUT("/api/categories/:categoryId", controller.Update)
	router.DELETE("/api/categories/:categoryId", controller.Delete)

	return router
}
