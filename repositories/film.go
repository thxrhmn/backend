package repositories

import (
  "week-02-task/models"
  "gorm.io/gorm"
)

type FilmRepository interface {
  FindFilms() ([]models.Films, error) // memasukan kedalam slice
  GetFilm(ID int) (models.Films, error)
	CreateFilm(film models.Films) (models.Films, error)
	UpdateFilm(film models.Films) (models.Films, error)
	DeleteFilm(film models.Films, ID int) (models.Films, error)
}

func RepositoryFilm(db *gorm.DB) *repository {
  return &repository{db}
}

func (r *repository) FindFilms() ([]models.Films, error) {
  var films []models.Films
  err := r.db.Preload("Category").Find(&films).Error // Using Find method ORM

  return films, err
}

func (r *repository) GetFilm(ID int) (models.Films, error) {
  var film models.Films
  err := r.db.Preload("Category").First(&film, ID).Error // Using First method ORM

  return film, err
}

func (r *repository) CreateFilm(film models.Films) (models.Films, error) {
  err := r.db.Create(&film).Error // Using Create method ORM

  return film, err
}

func (r *repository) UpdateFilm(film models.Films) (models.Films, error) {
  err := r.db.Save(&film).Error // Using Save method ORM

  return film, err
}

func (r *repository) DeleteFilm(film models.Films,ID int) (models.Films, error) {
	err := r.db.Delete(&film).Error // Using Delete method ORM

  return film, err
}