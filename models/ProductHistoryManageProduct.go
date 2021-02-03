package models

import "time"

type ProductHistoryManageProduct struct {
	Id                      int                  `json:"id" gorm:"primaryKey"`
	Product                 Product              `gorm:"foreignkey:Product_id;"`
	Product_id              int                  `json:"product_id" gorm:"not null;comment:สินค้า"`
	HistoryManageProduct    HistoryManageProduct `gorm:"foreignkey:HistoryManageProduct_id;"`
	HistoryManageProduct_id int                  `json:"history_manage_product_id" gorm:"not null;comment:ประวัติทำรายการ"`
	EditProduct             EditProduct          `gorm:"foreignkey:EditProduct_id;"`
	EditProduct_id          int                  `json:"edit_product_id" gorm:"not null;comment:ข้อมูลที่เปลี่ยนสินค้า"`
	Created_at              time.Time            `json:"created_at" gorm:"autoCreateTime"`
	Updated_at              time.Time            `json:"updated_at" gorm:"autoCreateTime"`
}
