package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Nama        string   `json:"nama" form:"nama"`
	Description string   `json:"description" form:"description"`
	Harga       int      `json:"harga" form:"harga"`
	CategoryID  int      `json:"category_id" form:"category_id"`
	Stock       int      `json:"stock" form:"stock"`
	Category    Category `json:"category"`
}
