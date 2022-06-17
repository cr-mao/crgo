package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Code  string `gorm:"column:code"`
	Price uint   `gorm:"column:price"`
}

// TableName maps to mysql table name.
func (p *Product) TableName() string {
	return "product"
}