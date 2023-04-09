package routes

import (
	"week-02-task/handlers"
	"week-02-task/pkg/middleware"
	"week-02-task/pkg/mysql"
	"week-02-task/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoute(e *echo.Group) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	e.GET("/transactions", h.FindTransactions)
	e.GET("/transaction/:id", h.GetTransaction)
	e.POST("/transaction", middleware.Auth(h.CreateTransaction))
	e.PATCH("/transaction/:id", middleware.Auth(h.UpdateTransaction))
	e.DELETE("/transaction/:id", middleware.Auth(h.DeleteTransaction))
}
