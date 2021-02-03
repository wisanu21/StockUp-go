package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_25_03_create_prefix_names() {
	name_file := "2021_01_25_03_create_prefix_names"

	if !IsHaveMigrations(name_file) {
		println(name_file, " ----> migrations loading")

		config.GetDB().Migrator().DropTable("prefix_names")
		config.GetDB().Migrator().CreateTable(&models.PrefixName{})
		SetDataMigrations(name_file)
		println(name_file, " ----> migrations successfully ")

	}
}
