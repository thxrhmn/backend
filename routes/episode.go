package routes

import (
	"week-02-task/handlers"
	"week-02-task/pkg/middleware"
	"week-02-task/pkg/mysql"
	"week-02-task/repositories"

	"github.com/labstack/echo/v4"
)

func EpisodeRoute(e *echo.Group) {
	episodeRepository := repositories.RepositoryEpisode(mysql.DB)
	h := handlers.HandlerEpisode(episodeRepository)

	e.GET("/film/:film_id/episodes", h.FindEpisodesByFilm)
	e.GET("/film/:film_id/episode/:episode_id", h.GetEpisodeByFilm)
	e.POST("/episode", middleware.Auth(middleware.UploadFile(h.CreateEpisode)))
	e.PATCH("/episode/:id", middleware.Auth(h.UpdateEpisode))
	e.DELETE("/episode/:id", middleware.Auth(h.DeleteEpisode))
}
