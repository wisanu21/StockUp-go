package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_25_05_create_menu() {
	name_file := "2021_01_25_05_create_menu"

	if !IsHaveMigrations(name_file) {
		println(name_file, " ----> migrations loading")

		config.GetDB().Migrator().DropTable("menus")
		config.GetDB().Migrator().CreateTable(&models.Menu{})
		SetDataMigrations(name_file)
		println(name_file, " ----> migrations successfully ")

	}
}
