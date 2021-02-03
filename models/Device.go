package models

import "time"

// "github.com/doug-martin/goqu/v9"

// "stockup-go/config"

type Device struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Employee    Employee  `gorm:"foreignkey:Employee_id;"`
	Employee_id int       `json:"employee_id" gorm:"not null"`
	Os          string    `json:"os"`
	Identifier  string    `json:"identifier"`
	Token       string    `json:"token"`
	Created_at  time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at  time.Time `json:"updated_at" gorm:"autoCreateTime"`
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
