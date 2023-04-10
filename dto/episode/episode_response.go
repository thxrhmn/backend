package episodedto

import "week-02-task/models"

type EpisodeResponse struct {
	Id            int          `json:"id"`
	Title         string       `json:"title" form:"title" validate:"required"`
	ThumbnailFilm string       `json:"thumbnailfilm" form:"thumbnailfilm" validate:"required"`
	LinkFilm      string       `json:"linkFilm" form:"linkFilm"`
	Film          models.Films `json:"film" form:"film"`
	FilmID        int          `json:"film_id" form:"film_id"`
}

type EpisodeDeleteResponse struct {
	ID int `json:"id"`
}
