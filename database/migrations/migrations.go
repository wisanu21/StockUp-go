package migrations

import (
	"stockup-go/config"
	"stockup-go/models"

	"github.com/gin-gonic/gin"
)

func RunMigrations(c *gin.Context) {
	// M_2021_01_25_01_create_companies()
	M_2021_01_25_01_create_migrations()
	M_2021_01_25_02_create_companies()
	M_2021_01_25_03_create_prefix_names()
	M_2021_01_25_04_create_levels()
	M_2021_01_25_05_create_menu()
	M_2021_01_25_06_create_employees()
	M_2021_01_25_07_create_device()
	M_2021_01_25_08_create_products()
	M_2021_01_26_01_create_history_manage_products()
	M_2021_01_26_02_create_menu_employees()
	M_2021_01_26_03_create_edit_products()
	M_2021_01_26_04_create_products_history_manage_products()

	// if errr.Error !=	 nil {
	// 	helper.Log(errr.Error.Error())
	// 	c.JSON(200, gin.H{"state": "error" + errr.Error.Error()})
	// 	return
	// } else {
	// 	c.JSON(200, gin.H{"state": "successfully"})
	// }
}

func IsHaveMigrations(table_name string) bool {
	// result := map[string]interface{}{}

	var migration models.Migration
	if config.GetDB().Migrator().HasTable(migration) {
		config.GetDB().Where("migration = ?", table_name).Find(&migration)
	} else {
		return false
	}

	if migration.Migration == "" {
		return false
	} else {
		return true
	}
}

func SetDataMigrations(table_name string) {
	var migration models.Migration
	migration.Migration = table_name
	config.GetDB().Save(&migration)
}
