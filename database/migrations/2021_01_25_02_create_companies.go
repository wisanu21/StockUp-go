package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_25_02_create_companies() {
	name_file := "2021_01_25_02_create_companies"

	if !IsHaveMigrations(name_file) {
		println("2021_01_25_02_create_companies ----> migrations loading")

		config.GetDB().Migrator().DropTable("companies")
		config.GetDB().Migrator().CreateTable(&models.Company{})
		SetDataMigrations(name_file)
		println("2021_01_25_02_create_companies ----> migrations successfully ")

	}
}
