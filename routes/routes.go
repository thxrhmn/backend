package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	UserRoute(e)
	FilmRoute(e)
	CategoryRoute(e)
	TransactionRoute(e)
	AuthRoutes(e)
	EpisodeRoute(e)
}
