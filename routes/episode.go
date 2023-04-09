package routes

import (
	"week-02-task/handlers"
	"week-02-task/pkg/mysql"
	"week-02-task/repositories"

	"github.com/labstack/echo/v4"
)

func EpisodeRoute(e *echo.Group) {
	episodeRepository := repositories.RepositoryEpisode(mysql.DB)
	h := handlers.HandlerEpisode(episodeRepository)

	e.GET("/episodes", h.FindEpisodes)
	e.GET("/episode/:id", h.GetEpisode)
	e.POST("/episode", h.CreateEpisode)
	e.PATCH("/episode/:id", h.UpdateEpisode)
	e.DELETE("/episode/:id", h.DeleteEpisode)
}
