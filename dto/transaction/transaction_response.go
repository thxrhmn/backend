package transactiondto

type TransactionResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type TransactionDeleteResponse struct {
	ID int `json:"id"`
}
