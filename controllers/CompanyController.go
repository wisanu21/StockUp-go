package Controller

import (
	"stockup-go/config"
	"stockup-go/models"

	"github.com/gin-gonic/gin"
)

func GetCompanyByIsRegister(c *gin.Context) {
	// helper.Log("GetCompanyByIsRegister()")
	var companyData []models.Company
	_ = config.GetDB().Where("is_register = ?", 1).Find(&companyData)
	c.JSON(200, companyData)
}
