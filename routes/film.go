package routes

import (
	"week-02-task/handlers"
	"week-02-task/pkg/middleware"
	"week-02-task/pkg/mysql"
	"week-02-task/repositories"

	"github.com/labstack/echo/v4"
)

func FilmRoute(e *echo.Group) {
	filmRepository := repositories.RepositoryFilm(mysql.DB)
	h := handlers.HandlerFilm(filmRepository)

	e.GET("/films", h.FindFilms)
	e.GET("/film/:id", h.GetFilm)
	e.POST("/film", middleware.UploadFile(h.CreateFilm))
	e.PATCH("/film/:id", middleware.Auth(h.UpdateFilm))
	e.DELETE("/film/:id", middleware.Auth(h.DeleteFilm))
}
