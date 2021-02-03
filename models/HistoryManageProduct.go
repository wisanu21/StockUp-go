package models

import "time"

type HistoryManageProduct struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Employee    Employee  `gorm:"foreignkey:Employee_id;"`
	Employee_id int       `json:"employee_id" gorm:"not null;comment:พนักงานที่ทำรายการเปลี่ยนแปลสินค้า"`
	Created_at  time.Time `json:"created_at" gorm:"autoCreateTime"`
	Updated_at  time.Time `json:"updated_at" gorm:"autoCreateTime"`
}
