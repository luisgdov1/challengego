package handlers

import (
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
	var jsondata map[string]string
	var user db.USER
	if err := c.ShouldBindJSON(&jsondata); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Where("rfc = ?", jsondata["rfc"]).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No existe registro"})
		return
	}
	rfc_value := jsondata["rfc"]
	file_name := utils.GenerateCSV(rfc_value)
	data_csv, erro := utils.ReadDataCSV(file_name[0])
	data_sending := utils.ClassifiedData(data_csv, erro)
	utils.Prepare_email(user.Name, user.Email, data_sending)
	c.JSON(http.StatusOK, gin.H{"data": jsondata})
}

func RendingEmailCSV(c *gin.Context) {
	data_csv, erro := utils.ReadDataCSV("test.csv")
	data_sending := utils.ClassifiedData(data_csv, erro)
	context := map[string]interface{}{
		"Name":                "Luis Gerardo",
		"Balance":             data_sending.Total_balance,
		"Promedio_Debito":     data_sending.Average_debit,
		"Prmedio_Credito":     data_sending.Average_credit,
		"Total_transacciones": data_sending.Total_transaction,
		"Operaciones":         data_sending.Transactions_per_month,
	}
	c.HTML(http.StatusOK, "free-simple-card.html", context)
}
func RendingEmailBD(c *gin.Context) {
	rfc := c.Param("rfc")
	clave_html := "free-simple-card.html"
	list_data := utils.GenerateCSV(rfc)
	file_name := string(list_data[0])
	data_csv, erro := utils.ReadDataCSV(file_name)
	if erro != nil {
		clave_html = "error.html"
	}
	data_sending := utils.ClassifiedData(data_csv, erro)
	context := map[string]interface{}{
		"Name":                string(list_data[1]),
		"Balance":             data_sending.Total_balance,
		"Promedio_Debito":     data_sending.Average_debit,
		"Prmedio_Credito":     data_sending.Average_credit,
		"Total_transacciones": data_sending.Total_transaction,
		"Operaciones":         data_sending.Transactions_per_month,
	}
	c.HTML(http.StatusOK, clave_html, context)
}
