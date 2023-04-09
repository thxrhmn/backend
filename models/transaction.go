package models

type Transactions struct {
	Id        int    `json:"id" gorm:"primary_key:auto_increment"`
	StartDate string `json:"startDate" gorm:"type: varchar(255)"`
	DueDate   string `json:"dueDate" gorm:"type: varchar(255)"`
	User      User   `json:"user"`
	UserID    int    `json:"user_id"`
	Attache   string `json:"attache" gorm:"type: varchar(255)"`
	Status    string `json:"status" gorm:"type: varchar(255)"`
}
