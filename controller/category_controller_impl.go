package controller

import (
	"go-restful/exception"
	"go-restful/helper"
	"go-restful/model/web"
	"go-restful/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	service service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		service: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadRequestBody(request, &categoryCreateRequest)

	categoryResponse, err := controller.service.Save(request.Context(), categoryCreateRequest)
	if err != nil {
		exception.HandleErrorRequest(err, writer)
		return
	}

	response := web.Response{
		Ok:      true,
		Code:    201,
		Message: "category added successfully",
		Data:    categoryResponse,
	}

	helper.WriteResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadRequestBody(request, &categoryUpdateRequest)

	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = categoryId
	categoryResponse, err := controller.service.Update(request.Context(), categoryUpdateRequest)
	if err != nil {
		exception.HandleErrorRequest(err, writer)
		return
	}

	response := web.Response{
		Ok:      true,
		Code:    200,
		Message: "category updated successfully",
		Data:    categoryResponse,
	}

	helper.WriteResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	err = controller.service.Delete(request.Context(), categoryId)
	if err != nil {
		exception.HandleErrorRequest(err, writer)
		return
	}

	response := web.Response{
		Ok:      true,
		Code:    200,
		Message: "category deleted successfully",
	}

	helper.WriteResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categories := controller.service.FindAll(request.Context())
	response := web.Response{
		Ok:      true,
		Code:    200,
		Message: "success get categories",
		Data:    categories,
	}

	helper.WriteResponseBody(writer, response)
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("categoryId"))
	helper.PanicIfError(err)

	category, err := controller.service.FindById(request.Context(), categoryId)
	if err != nil {
		exception.HandleErrorRequest(err, writer)
		return
	}

	response := web.Response{
		Ok:      true,
		Code:    200,
		Message: "success get category",
		Data:    category,
	}

	helper.WriteResponseBody(writer, response)
}
