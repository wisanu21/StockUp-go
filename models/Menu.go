package models

import "time"

type Menu struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"comment:ชื่อเมนู"`
	Is_active  bool    `json:"is_active" gorm:"comment:เปิดใช้โดยผู้พัฒนา"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updated_at" gorm:"autoCreateTime"`
	Deleted_at time.Time `json:"deleted_at" gorm:"default:null"`
}
