package models

import "time"

type Level struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name" gorm:"comment:ชื่อ"`
	Created_at time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at time.Time `json:"updated_at" gorm:"autoCreateTime"`
	Deleted_at time.Time `json:"deleted_at" gorm:"default:null"`
}
