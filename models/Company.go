package models

import "time"

type Company struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"comment:ชื่อ companies"`
	Is_active   bool      `json:"is_active" gorm:"comment:เปิดใช้โรงงาน"`
	Created_at  time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at  time.Time `json:"updated_at" gorm:"autoCreateTime"`
	Deleted_at  time.Time `json:"deleted_at" gorm:"default:null"`
	Is_register bool      `json:"is_register" gorm:"comment:แสดงใน select Register หรือไม่"`
	Path_image  string    `json:"path_image" gorm:"comment:ลิ้งรูป"`
}
