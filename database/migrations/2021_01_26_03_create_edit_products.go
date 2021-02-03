package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_26_03_create_edit_products() {
	name_file := "2021_01_26_03_create_edit_products"

	if !IsHaveMigrations(name_file) {
		println(name_file, " ----> migrations loading")

		config.GetDB().Migrator().DropTable("edit_products")
		config.GetDB().Migrator().CreateTable(&models.EditProduct{})
		SetDataMigrations(name_file)
		println(name_file, " ----> migrations successfully ")

	}
}
