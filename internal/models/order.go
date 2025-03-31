package models

import "time"

type Order struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"product_id" gorm:"not null"`
	Product      Product `json:"product" gorm:"foreignKey:ProductRefer"`
	UserRefer    int     `json:"user_id" gorm:"not null"`
	User         User    `json:"user" gorm:"foreignKey:UserRefer"`
}
