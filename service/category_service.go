package service

import (
	"context"
	"go-restful/model/web"
)

type CategoryService interface {
	Save(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error)
	Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error)
	Delete(ctx context.Context, categoryId int) error
	FindAll(ctx context.Context) []web.CategoryResponse
	FindById(ctx context.Context, categoryId int) (web.CategoryResponse, error)
}
