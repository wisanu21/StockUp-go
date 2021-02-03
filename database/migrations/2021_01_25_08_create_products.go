package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_25_08_create_products() {
	name_file := "2021_01_25_08_create_products"

	if !IsHaveMigrations(name_file) {
		println(name_file, " ----> migrations loading")

		config.GetDB().Migrator().DropTable("products")
		config.GetDB().Migrator().CreateTable(&models.Product{})
		SetDataMigrations(name_file)
		println(name_file, " ----> migrations successfully ")

	}
}
