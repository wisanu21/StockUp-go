package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_26_04_create_products_history_manage_products() {
	name_file := "2021_01_26_04_create_products_history_manage_products"

	if !IsHaveMigrations(name_file) {
		println(name_file, " ----> migrations loading")

		config.GetDB().Migrator().DropTable("products_history_manage_products")
		config.GetDB().Migrator().CreateTable(&models.ProductHistoryManageProduct{})
		SetDataMigrations(name_file)
		println(name_file, " ----> migrations successfully ")

	}
}
