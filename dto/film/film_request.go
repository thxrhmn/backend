package filmdto

import "week-02-task/models"

type CreateFilmRequest struct {
	Id            int             `json:"id"`
	Title         string          `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string          `json:"thumbnailfilm" form:"thumbnailfilm" validate:"required"`
	Year          int             `json:"year" form:"year" validate:"required"`
	Category      models.Category `json:"category" form:"category" validate:"required"`
	CategoryID    int             `json:"category_id" form:"category_id" validate:"required"`
	Description   string          `json:"description" form:"description" validate:"required"`
}

type UpdateFilmRequest struct {
	Id            int             `json:"id"`
	Title         string          `json:"title" form:"title"`
	ThumbnailFilm string          `json:"thumbnailfilm" form:"thumbnailfilm"`
	Year          int             `json:"year" form:"year"`
	Category      models.Category `json:"category" form:"category" validate:"required"`
	CategoryID    int             `json:"category_id" form:"category_id" validate:"required"`
	Description   string          `json:"description" form:"description"`
}
