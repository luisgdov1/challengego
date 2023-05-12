package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/challengego/db"
	"github.com/challengego/handlers"
)

func main() {
	fmt.Println("Hola go")
	db.ConnectDB()
	router := gin.Default()
	router.GET("/balance", handlers.GetResumenCSV)
	router.Run(":8000")
}
