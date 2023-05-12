package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RESUMEN struct {
	Total_balance          float32        `json:"total_balance"`
	Total_transaction      int            `json:"total_transaction"`
	Average_debit          float32        `json:"average_debit"`
	Average_credit         float32        `json:"average_credit"`
	Transactions_per_month []TRANSACTIONS_RESUMEN `json:"transactions"`
}

type TRANSACTIONS_RESUMEN struct {
	Month               string `json:"month"`
	Number_transactions int    `json:"number_transactions"`
}

type USER struct {
	gorm.Model
	RFC      string `json:"rfc" gorm:"primary_key"`
	Name     string `json:"nombre"`
	LastName string `json:"apellido"`
}

type VISITLOG struct {
	gorm.Model
	DateVisit time.Time `json:"date_visit"`
	RFC       USER      `json:"RFC" gorm:"foreignKey:USER"`
}

type OPERATION struct{
	gorm.Model
	RFC USER `json:"RFC" gorm:"foreignKey:USER"`
	Type_Operation string `json:"Operation_Type"`
	Balance string `json:"Balance_Operation"`
}

func ConnectDB()  {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&USER{})
	db.AutoMigrate(&VISITLOG{})
	db.AutoMigrate(&OPERATION{})
}