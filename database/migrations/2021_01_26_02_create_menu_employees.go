package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_26_02_create_menu_employees() {
	name_file := "2021_01_26_02_create_menu_employees"

	if !IsHaveMigrations(name_file) {
		println(name_file, " ----> migrations loading")

		config.GetDB().Migrator().DropTable("menu_employees")
		config.GetDB().Migrator().CreateTable(&models.MenuEmployee{})
		SetDataMigrations(name_file)
		println(name_file, " ----> migrations successfully ")

	}
}
