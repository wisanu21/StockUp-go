package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_25_07_create_device() {
	name_file := "2021_01_25_07_create_device"

	if !IsHaveMigrations(name_file) {
		println(name_file, " ----> migrations loading")

		config.GetDB().Migrator().DropTable("devices")
		config.GetDB().Migrator().CreateTable(&models.Device{})
		SetDataMigrations(name_file)
		println(name_file, " ----> migrations successfully ")

	}
}
