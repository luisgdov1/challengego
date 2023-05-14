package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/challengego/db"
	"github.com/challengego/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("ERROR EN EL .ENV")
		fmt.Print(err)
	}
	db.ConnectDB()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("templates/", "templates/")
	router.GET("/balance", handlers.GetResumenCSV)
	router.POST("/createUser", handlers.CreateUser)
	router.POST("/createOperation", handlers.CreateOperation)
	router.GET("/allOperation", handlers.GetOperations)
	router.POST("/sendingEmail", handlers.SendingEmailBD)
	router.GET("/emailPreview", handlers.RendingEmailCSV)
	router.GET("/email/:rfc", handlers.RendingEmailBD)
	router.Run(":8000")
}
