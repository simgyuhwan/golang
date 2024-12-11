package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"index:unique"`
	Email string `json:"email"`
	Password string `json:"-" gorm:"column:password"`
	Address string `json:"address"`
	Phone string `json:"phone"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime:nano"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}