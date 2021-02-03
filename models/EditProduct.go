package models

import (
	"time"
)

type EditProduct struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"comment:ชื่อ"`
	Price       float64   `json:"price" gorm:"comment:ราคา"`
	Image_path  string    `json:"image_path" gorm:"comment:ที่อยู่รูป"`
	Num_product int       `json:"num_product" gorm:"comment:จำนวนสินค้า"`
	Detail      string    `json:"detail" gorm:"comment:ข้อมูล"`
	Created_at  time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at  time.Time `json:"updated_at" gorm:"autoCreateTime"`
	Deleted_at  time.Time `json:"deleted_at" gorm:"default:null"`
}
