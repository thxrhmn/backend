package transactiondto

import "week-02-task/models"

type CreateTransactionRequest struct {
	Id        int         `json:"id"`
	StartDate string      `json:"startDate"`
	DueDate   string      `json:"dueDate"`
	User      models.User `json:"user"`
	UserID    int         `json:"user_id"`
	Attache   string      `json:"attache"`
	Status    string      `json:"status"`
}

type UpdateTransactionRequest struct {
	Id        int         `json:"id"`
	StartDate string      `json:"startDate"`
	DueDate   string      `json:"dueDate"`
	User      models.User `json:"user"`
	UserID    int         `json:"user_id"`
	Attache   string      `json:"attache"`
	Status    string      `json:"status"`
}
