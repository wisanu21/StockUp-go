package migrations

import (
	"stockup-go/config"
	"stockup-go/models"
)

func M_2021_01_25_01_create_migrations() {
	name_file := "2021_01_25_01_create_migrations"

	if !IsHaveMigrations(name_file) {
		println("M_2021_01_25_01_create_migrations ----> migrations loading")
		config.GetDB().Migrator().DropTable("migrations")
		config.GetDB().Migrator().CreateTable(&models.Migration{})
		SetDataMigrations(name_file)
		println("M_2021_01_25_01_create_migrations ----> migrations successfully ")
	}

}

