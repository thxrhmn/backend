package episodedto

import "week-02-task/models"

type CreateEpisodeRequest struct {
	Id            int          `json:"id"`
	Title         string       `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string       `json:"thumbnailfilm" form:"thumbnailfilm" validate:"required"`
	LinkFilm      string       `json:"linkFilm" form:"year" validate:"required"`
	Film          models.Films `json:"film"`
	FilmID        int          `json:"film_id"`
}

type UpdateEpisodeRequest struct {
	Id            int          `json:"id"`
	Title         string       `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string       `json:"thumbnailfilm" form:"thumbnailfilm" validate:"required"`
	LinkFilm      string       `json:"linkFilm" form:"year" validate:"required"`
	Film          models.Films `json:"film"`
	FilmID        int          `json:"film_id"`
}
