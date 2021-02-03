package models

import "time"

// "github.com/doug-martin/goqu/v9"

// "stockup-go/config"

type Employee struct {
	Id             int        `json:"id" gorm:"primaryKey"`
	First_name     string     `json:"first_name"`
	Last_name      string     `json:"last_name"`
	Mobile         string     `json:"mobile"`
	Password       string     `json:"password"`
	Companies      Company    `gorm:"foreignkey:Company_id;"`
	Company_id     int        `json:"company_id" gorm:"not null"`
	Prefix_names   PrefixName `gorm:"foreignKey:Prefix_name_id;"`
	Prefix_name_id int        `json:"prefix_name_id" gorm:"default:null"`
	Created_at     time.Time  `json:"created_at" gorm:"autoCreateTime"`
	Updated_at     time.Time  `json:"updated_at" gorm:"autoCreateTime"`
	Deleted_at     time.Time  `json:"deleted_at" gorm:"default:null"`
	Easy_name      string     `json:"easy_name"`
	Is_active      int        `json:"is_active"`
	Levels         Level      `gorm:"foreignKey:Level_id;"`
	Level_id       int        `json:"level_id" gorm:"default:null"`
}

// type DeviceModel struct {
// 	Table_name string "devices"
// }

// var Table_name = "devices"

// type DeviceModel struct {
// }

// func (a DeviceModel) First(id int) int {
// 	id = 123
// 	return id
// }
