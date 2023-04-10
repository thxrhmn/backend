package repositories

import (
	"week-02-task/models"

	"gorm.io/gorm"
)

type EpisodeRepository interface {
	FindEpisodes() ([]models.Episode, error) // memasukan kedalam slice
	GetEpisode(ID int) (models.Episode, error)
	FindEpisodesByFilm(ID int) ([]models.Episode, error)
	GetEpisodeByFilm(ID int, Id int) (models.Episode, error)
	CreateEpisode(episode models.Episode) (models.Episode, error)
	UpdateEpisode(episode models.Episode) (models.Episode, error)
	DeleteEpisode(episode models.Episode, ID int) (models.Episode, error)
}

func RepositoryEpisode(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindEpisodesByFilm(ID int) ([]models.Episode, error) {
	var episode []models.Episode
	err := r.db.Preload("Film").Where("film_id = ?", ID).Find(&episode).Error

	return episode, err
}

func (r *repository) GetEpisodeByFilm(ID int, EpisodeID int) (models.Episode, error) {
	var episode models.Episode
	err := r.db.Preload("Film").Where("film_id = ? AND id = ?", ID, EpisodeID).First(&episode).Error

	return episode, err
}

func (r *repository) FindEpisodes() ([]models.Episode, error) {
	var episode []models.Episode
	err := r.db.Preload("Film").Find(&episode).Error // Using Find method ORM

	return episode, err
}

func (r *repository) GetEpisode(ID int) (models.Episode, error) {
	var episode models.Episode
	err := r.db.Preload("Film").First(&episode, ID).Error // Using First method ORM

	return episode, err
}

func (r *repository) CreateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Preload("Film").Create(&episode).Error // Using Create method ORM

	return episode, err
}

func (r *repository) UpdateEpisode(episode models.Episode) (models.Episode, error) {
	err := r.db.Save(&episode).Error // Using Save method ORM

	return episode, err
}

func (r *repository) DeleteEpisode(episode models.Episode, ID int) (models.Episode, error) {
	err := r.db.Delete(&episode).Error // Using Delete method ORM

	return episode, err
}
