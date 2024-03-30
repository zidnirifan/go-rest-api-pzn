package web

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required,min=1,max=200" json:"name"`
}

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Response struct {
	Ok      bool        `json:"ok"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
