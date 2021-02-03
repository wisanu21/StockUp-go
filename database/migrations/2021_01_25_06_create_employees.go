package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_25_06_create_employees() {
	name_file := "2021_01_25_06_create_employees"

	if !IsHaveMigrations(name_file) {
		println(name_file, " ----> migrations loading")

		config.GetDB().Migrator().DropTable("employees")
		config.GetDB().Migrator().CreateTable(&models.Employee{})
		SetDataMigrations(name_file)
		println(name_file, " ----> migrations successfully ")

	}
}
