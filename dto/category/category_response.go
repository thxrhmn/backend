package categorydto

type CategoryResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CategoryDeleteResponse struct {
	ID int `json:"id"`
}
