package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/challengego/utils"
)

func GetResumenCSV(c *gin.Context) {
	data_csv, erro := utils.ReadDataCSV("test.csv")
	data_sending := utils.ClassifiedData(data_csv, erro)
	c.JSON(http.StatusOK, data_sending)
}
