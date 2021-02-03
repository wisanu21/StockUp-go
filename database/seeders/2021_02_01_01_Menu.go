package seeders

import (
	"stockup-go/config"
	"stockup-go/helper"
	"stockup-go/models"
)

func S_2021_02_01_01_Menu() {
	name_file := "S_2021_02_01_01_Menu"
	println(name_file, " ---->seeders loading")
	// arrMenus := [][]string{{"จัดการพนักงาน","1"}, {"จัดการร้านค้า","1"}, {"จัดการสินค้า","1"}, {"จัดการข้อมูลส่วนตัว","1"}}

	arrMenus := []struct {
		name      string
		is_active bool
	}{
		{"จัดการพนักงาน", true},
		{"จัดการร้านค้า", true},
		{"จัดการสินค้า", true},
		{"จัดการข้อมูลส่วนตัว", true},
	}

	print(len(arrMenus))

	for id := 1; id <= len(arrMenus); id++ {
		var menus models.Menu
		menus.Id = id
		menus.Name = arrMenus[id-1].name
		menus.Is_active = arrMenus[id-1].is_active
		var err = config.GetDB().Save(&menus)
		if err.Error != nil {
			helper.Log(err.Error.Error())
		} else {
			println(name_file, " : ", arrMenus[id-1].name, " ----> seeders No Error")
		}
	}
	println(name_file, " ----> seeders successfully ")
}
