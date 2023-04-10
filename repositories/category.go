package repositories

import (
	"week-02-task/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindCategories() ([]models.Category, error) // memasukan kedalam slice
	GetCategory(ID int) (models.Category, error)
	CreateCategory(category models.Category) (models.Category, error)
	UpdateCategory(category models.Category) (models.Category, error)
	DeleteCategory(category models.Category, ID int) (models.Category, error)
}

func RepositoryCategory(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error // Using Find method ORM

	return categories, err
}

func (r *repository) GetCategory(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.First(&category, ID).Error // Using First method ORM

	return category, err
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error // Using Create method ORM

	return category, err
}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error // Using Save method ORM

	return category, err
}

func (r *repository) DeleteCategory(category models.Category, ID int) (models.Category, error) {
	err := r.db.Delete(&category).Error // Using Delete method ORM

	return category, err
}
