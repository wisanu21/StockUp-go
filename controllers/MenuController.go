package Controller

import (
	"stockup-go/config"
	"stockup-go/models"

	"github.com/gin-gonic/gin"
)

func GetMenuByIsActive(c *gin.Context) {
	// helper.Log("GetCompanyByIsRegister()")
	var memu []models.Menu
	_ = config.GetDB().Where("is_active = ?", 1).Find(&memu)
	c.JSON(200, memu)
}
