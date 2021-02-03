package Controller

import (
	"stockup-go/config"
	"stockup-go/models"
	"strconv"
	"time"

	helper "stockup-go/helper"

	"github.com/gin-gonic/gin"
)

// "go-api/handle"

type LoginController struct{}

type request_LoginByDeviceID struct {
	Os         string `json:"os"`
	Identifier string `json:"identifier"`
}

func LoginByDeviceID(c *gin.Context) {
	helper.Log("LoginByDeviceID()")
	var request request_LoginByDeviceID
	c.BindJSON(&request)
	var deviceData models.Device
	err := config.GetDB().Where("identifier = ?", request.Identifier).Where("os = ?", request.Os).First(&deviceData)

	if err.Error != nil {
		helper.Log(err.Error.Error())
		c.JSON(200, gin.H{"state": "error"})
		return
	} else {
		helper.Log(helper.Concatenation("login by device , ID in table is ", strconv.Itoa(deviceData.Id)))
		deviceData.Token = helper.RandStringRunes(25)
		deviceData.Updated_at = time.Now()
		helper.Log(helper.Concatenation("new device token for login by device is", deviceData.Token))
		config.GetDB().Save(&deviceData)
		helper.Log("update token device successfully")
		c.JSON(200, gin.H{"state": "successfully", "detail": "login", "employee_id": deviceData.Employee_id, "token": deviceData.Token})
	}

}

type request_Login struct {
	Os         string `json:"os"`
	Identifier string `json:"identifier"`
	Phone      string `json:"text_phone"`
	Password   string `json:"text_password"`
}

func Login(c *gin.Context) {
	helper.Log("Login()")
	var request request_Login
	c.BindJSON(&request)
	var employeeData models.Employee

	err := config.GetDB().Where("mobile = ?", request.Phone).Where("password = ?", request.Password).First(&employeeData)

	if err.Error != nil {
		helper.Log(err.Error.Error())
		c.JSON(200, gin.H{"error": "failed"})
		return
	} else {
		helper.Log(helper.Concatenation("login by mobile password , ID in table is ", strconv.Itoa(employeeData.Id)))
		var deviceData models.Device
		deviceData.Employee_id = employeeData.Id
		deviceData.Os = request.Os
		deviceData.Identifier = request.Identifier
		deviceData.Created_at = time.Now()
		deviceData.Updated_at = time.Now()
		deviceData.Token = helper.RandStringRunes(25)
		config.GetDB().Save(&deviceData)
		helper.Log(helper.Concatenation("save device , table device id is : ", strconv.Itoa(deviceData.Id)))
		// c.JSON(200, gin.H{"state": "successfully"})
		c.JSON(200, gin.H{"state": "successfully", "detail": "login", "employee_id": deviceData.Employee_id, "token": deviceData.Token})
	}
}

type request_Logout struct {
	Os         string `json:"os"`
	Identifier string `json:"identifier"`
}

func Logout(c *gin.Context) {
	helper.Log("Logout()")
	var request request_Logout
	c.BindJSON(&request)
	var deviceData models.Device
	err := config.GetDB().Where("identifier = ?", request.Identifier).Where("os = ?", request.Os).Delete(&deviceData)
	if err.Error != nil {
		helper.Log(err.Error.Error())
		c.JSON(200, gin.H{"error": "failed"})
		return
	} else {
		helper.Log(helper.Concatenation("Logout , employee ID in table is ", strconv.Itoa(deviceData.Employee_id)))
		c.JSON(200, gin.H{"state": "successfully"})
	}
}
