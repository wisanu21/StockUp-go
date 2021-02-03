package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_25_04_create_levels() {
	name_file := "2021_01_25_04_create_levels"

	if !IsHaveMigrations(name_file) {
		println(name_file, " ----> migrations loading")

		config.GetDB().Migrator().DropTable("levels")
		config.GetDB().Migrator().CreateTable(&models.Level{})
		SetDataMigrations(name_file)
		println(name_file, " ----> migrations successfully ")

	}
}
