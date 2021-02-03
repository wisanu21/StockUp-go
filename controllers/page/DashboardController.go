package Controller

import (
	"fmt"
	"stockup-go/config"
	"stockup-go/helper"

	"github.com/gin-gonic/gin"
)

// type request_GetDataDashboard struct {
// 	Token       string `json:"token"`
// 	Employee_id string `json:"employee_id" `
// }

func GetDataDashboard(c *gin.Context) {
	helper.Log("GetDataDashboard()")
	request := map[string]interface{}{}
	c.BindJSON(&request)
	// fmt.Println(request["employee_id"])

	result := map[string]interface{}{}
	err := config.GetDB().
		Table("employees").
		Select(
			"employees.first_name as employees_first_name," +
				"employees.last_name as employees_last_name," +
				"levels.name as levels_name," +
				"employees.id as employees_id" +
				"").
		// Joins("left join companies on companies.id = employees.company_id").
		Joins("left join levels on levels.id = employees.level_id").
		Find(&result)
	fmt.Println(result)
	if err.Error != nil {
		helper.Log(err.Error.Error())
		c.JSON(200, gin.H{"state": "error"})
		return
	} else {
		// var menu []string
		// menu[0] = "sss"
		result["arr_name_menu"], result["arr_id_menu"] = GetSidebarByEmployeeid(result["employees_id"])
		helper.Log("successfully GetDataDashboard")

		c.JSON(200, result)
	}
	// println(result)
	// c.JSON(200, result)
}
