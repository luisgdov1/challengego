package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/challengego/db"
	"github.com/challengego/utils"
)

func GetResumenCSV(c *gin.Context) {
	data_csv, erro := utils.ReadDataCSV("test.csv")
	data_sending := utils.ClassifiedData(data_csv, erro)
	c.JSON(http.StatusOK, data_sending)
}

func CreateUser(c *gin.Context) {
	var user db.USER

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func CreateOperation(c *gin.Context) {
	var operation db.OPERATION
	var user db.USER
	if err := c.ShouldBindJSON(&operation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var_object := operation.RFC
	db.DB.Where("rfc = ?", var_object).First(&user)
	operation.User = user
	db.DB.Create(&operation)
	c.JSON(http.StatusOK, gin.H{"data": operation})
}

func GetOperations(c *gin.Context) {
	var operation []db.OPERATION
	db.DB.Find(&operation)
	c.JSON(http.StatusOK, gin.H{"data": operation})
}

func SendingEmailBD(c *gin.Context) {
	var jsondata map[string]interface{}
	if err := c.ShouldBindJSON(&jsondata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(jsondata)
	c.JSON(http.StatusOK, gin.H{"data": jsondata})
}
