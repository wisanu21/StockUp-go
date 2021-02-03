package models

import "time"

type MenuEmployee struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Employee    Employee  `gorm:"foreignkey:Employee_id;"`
	Employee_id int       `json:"employee_id" gorm:"not null;comment:พนักงาน"`
	Menu    Menu  `gorm:"foreignkey:Menu_id;"`
	Menu_id int       `json:"menu_id" gorm:"not null;comment:เมนู"`
	Created_at  time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at  time.Time `json:"updated_at" gorm:"autoCreateTime"`
}
