package database

import (
	"fmt"
	"week-02-task/models"
	"week-02-task/pkg/mysql"
)

// Automatic Migration if Running App
// otomatis membuat table ketika memasukan models keadalam method AutoMigrate()
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Films{},
		&models.Category{},
		&models.Episode{},
		&models.Transactions{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
