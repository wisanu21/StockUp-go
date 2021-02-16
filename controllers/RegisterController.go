package Controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"stockup-go/config"
	helper "stockup-go/helper"
	"stockup-go/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterController struct{}

type request_SaveRegister struct {
	Select_company   string `json:"select_company"`
	Text_fname       string `json:"text_fname"`
	Text_lname       string `json:"text_lname"`
	Text_phone       string `json:"text_phone"`
	Text_password    string `json:"text_password"`
	Text_easy_name   string `json:"text_easy_name"`
	Bytes_image_file string `json:"bytes_image_file"`
}

func SaveRegister(c *gin.Context) {
	// print("sdf")
	helper.Log("SaveRegister()")

	var request request_SaveRegister
	c.BindJSON(&request)
	var employeesData []models.Employee
	err := config.GetDB().Where("company_id = ?", request.Select_company).Find(&employeesData)
	// print(len(employeesData), request.Select_company)
	if err.Error != nil {
		helper.Log(err.Error.Error())
		c.JSON(200, gin.H{"error": "failed"})
		return
	} else {
		helper.Log("add data in employeeData")
		var employeeData models.Employee
		if len(employeesData) == 0 {
			employeeData.Level_id = 1
			employeeData.Is_active = 1
		} else {
			// employeeData.Level_id = is null
			employeeData.Is_active = 0
		}
		// error errr ;

		employeeData.First_name = request.Text_fname
		employeeData.Last_name = request.Text_lname
		employeeData.Mobile = request.Text_phone
		employeeData.Password = request.Text_password

		Company_id, _ := strconv.Atoi(request.Select_company)
		employeeData.Company_id = Company_id
		employeeData.Updated_at = time.Now()
		employeeData.Created_at = time.Now()
		employeeData.Easy_name = request.Text_easy_name
		// employeeData.Prefix_name_id =
		helper.Log("end save detail employeeData")
		helper.Log("save image employeeData")
		employeeData.Path_image = saveImageEmployee(request.Bytes_image_file)
		helper.Log("befor save employeeData")
		err = config.GetDB().Save(&employeeData)
		fmt.Println(employeeData.Id)
		helper.Log("after save employeeData")
		if employeeData.Level_id == 1 {
			addMenuFirstEmployee(employeeData.Id)
		}

		// print(&(err.Error), "------")
		// print(err, "dsf")
		if err.Error != nil {
			helper.Log(err.Error.Error())
			c.JSON(200, gin.H{"state": "error", "detail": err.Error.Error()})
		} else {
			helper.Log("no err")
			if employeeData.Is_active == 1 {
				c.JSON(200, gin.H{"state": "successfully", "detail": "สามารถเข้าสู่ระบบได้"})
			} else {
				c.JSON(200, gin.H{"state": "successfully", "detail": "รอเจ้าของร้านอนุญาต จึงสามารถเข้าสู่ระบบ"})
			}
		}

	}

}

func saveImageEmployee(bytes_image_file string) string {
	// return "dsdf"
	coI := strings.Index(string(bytes_image_file), ",")
	rawImage := string(bytes_image_file)[coI+1:]
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))
	res := bytes.NewReader(unbased)

	file, err := os.Create("storage/employee/asdf.jpg")
	///storage/employee/asdf.jpg
	if err != nil {
		helper.Log(err.Error())
	}
	defer file.Close()
	io.Copy(file, res)

	return " "
}

func addMenuFirstEmployee(employee_id int) {
	helper.Log(helper.Concatenation("addMenuFirstEmployee()", strconv.Itoa(employee_id)))
	var menu []models.Menu
	err := config.GetDB().Where("is_active = ?", 1).Find(&menu)
	if err.Error != nil {
		helper.Log("Err get menu")
		helper.Log(err.Error.Error())
	} else {
		for i := 0; i < len(menu); i++ {
			var menuEmployee models.MenuEmployee
			menuEmployee.Employee_id = employee_id
			menuEmployee.Menu_id = menu[i].Id
			errr := config.GetDB().Save(&menuEmployee)
			// helper.Log(helper.Concatenation(" addMenuFirstEmployee()", strconv.Itoa(employee_id)))
			if errr.Error != nil {
				helper.Log("Err save menuEmployee")
				helper.Log(errr.Error.Error())
			} else {
				helper.Log("save menuEmployee ")
				helper.Log(helper.Concatenation("employee_id : ", strconv.Itoa(employee_id)))
				helper.Log(helper.Concatenation("Menu_id : ", strconv.Itoa(menu[i].Id)))
			}
		}
	}
}
