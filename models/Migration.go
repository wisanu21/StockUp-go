package models

// "github.com/doug-martin/goqu/v9"

// "stockup-go/config"

type Migration struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Migration string `json:"migration" gorm:"comment:ชื่อ migration"`
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
