package Controller

import (
	"fmt"
	"stockup-go/config"
	"stockup-go/helper"
	"stockup-go/models"
)

func GetSidebarByEmployeeid(employee_id interface{}) ([]string, []int) {
	helper.Log("GetSidebarByEmployeeid()")
	fmt.Println(employee_id)
	var menuEmployee []models.MenuEmployee
	err := config.GetDB().Where("employee_id = ?", employee_id).Find(&menuEmployee)
	var arr_name_menu []string
	var arr_id_menu []int
	// var arr_menu [][]int

	if err.Error != nil {
		helper.Log("Err get menuEmployee")
		helper.Log(err.Error.Error())
	} else {
		fmt.Println(len(menuEmployee))

		for i := 0; i < len(menuEmployee); i++ {
			fmt.Println(menuEmployee[i].Menu_id)
			var menu models.Menu
			errr := config.GetDB().Where("is_active = ?", 1).Where("id = ?", menuEmployee[i].Menu_id).First(&menu)
			if errr.Error != nil {
				helper.Log("Err get menu")
				helper.Log(errr.Error.Error())
			} else {
				fmt.Println(menu.Name)
				arr_name_menu = append(arr_name_menu, menu.Name)
				arr_id_menu = append(arr_id_menu, menu.Id)
			}
		}

		// fmt.Println("2d: ", arr_menu)

	}
	// fmt.Printf("%v", arr_menu)
	return arr_name_menu, arr_id_menu
}
