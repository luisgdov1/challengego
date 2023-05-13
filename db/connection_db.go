package db

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RESUMEN struct {
	Total_balance          float32                `json:"total_balance"`
	Total_transaction      int                    `json:"total_transaction"`
	Average_debit          float32                `json:"average_debit"`
	Average_credit         float32                `json:"average_credit"`
	Transactions_per_month []TRANSACTIONS_RESUMEN `json:"transactions"`
}

type TRANSACTIONS_RESUMEN struct {
	Month               string `json:"month"`
	Number_transactions int    `json:"number_transactions"`
}

type OPERATION struct {
	gorm.Model
	ID             uint `json:"id_operation" gorm:"autoIncrement:true"`
	UserID         uint
	User           USER      `gorm:"foreignKey:UserID"`
	Type_Operation string    `json:"Type_Operation"`
	Balance        float32   `json:"Balance"`
	DateVisit      time.Time `json:"DateVisit"`
	RFC            string    `json:"RFC"`
}

type USER struct {
	gorm.Model
	ID       uint   `gorm:"primary_key;autoIncrement:true"`
	RFC      string `json:"RFC"`
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	Email    string `json:"Email"`
}

var DB *gorm.DB

func (operation *OPERATION) BeforeCreate(tx *gorm.DB) (err error) {
	operation.DateVisit = time.Now()
	return nil
}

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&USER{})
	db.AutoMigrate(&OPERATION{})
	DB = db
}
