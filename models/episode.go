package models

type Episode struct {
	Id            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailFilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	LinkFilm      string `json:"linkFilm" form:"linkFilm" gorm:"type: varchar(255)"`
	Film          Films  `json:"film"`
	FilmID        int    `json:"film_id"`
}
