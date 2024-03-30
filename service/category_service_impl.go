package service

import (
	"context"
	"database/sql"
	"go-restful/exception"
	"go-restful/helper"
	"go-restful/model/domain"
	"go-restful/model/web"
	"go-restful/repository"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *CategoryServiceImpl) Save(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.CategoryResponse{}, err
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.Repository.Save(ctx, tx, category)

	return web.CategoryResponse(category), nil
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.CategoryResponse{}, err
	}

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, request.Id)
	if err != nil {
		return web.CategoryResponse{}, &exception.RequestError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}
	category = service.Repository.Update(ctx, tx, domain.Category(request))

	return web.CategoryResponse(category), nil
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.Repository.FindById(ctx, tx, categoryId)
	if err != nil {
		return &exception.RequestError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	service.Repository.Delete(ctx, tx, categoryId)
	return nil
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.Repository.FindAll(ctx, tx)

	var categoriesResponse []web.CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, web.CategoryResponse(category))
	}

	return categoriesResponse
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) (web.CategoryResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId)
	if err != nil {
		return web.CategoryResponse{}, &exception.RequestError{
			StatusCode: http.StatusNotFound,
			Err:        err,
		}
	}

	return web.CategoryResponse(category), nil
}
