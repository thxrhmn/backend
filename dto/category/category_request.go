package categorydto

type CreateCategoryRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UpdateCategoryRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
