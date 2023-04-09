package routes

import (
	"week-02-task/handlers"
	"week-02-task/pkg/middleware"
	"week-02-task/pkg/mysql"
	"week-02-task/repositories"

	"github.com/labstack/echo/v4"
)

func CategoryRoute(e *echo.Group) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	e.GET("/categories", h.FindCategories)
	e.GET("/category/:id", h.GetCategory)
	e.POST("/category", middleware.Auth(h.CreateCategory))
	e.PATCH("/category/:id", middleware.Auth(h.UpdateCategory))
	e.DELETE("/category/:id", middleware.Auth(h.DeleteCategory))
}
