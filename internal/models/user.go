package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null;default:null"`
	Email    string `gorm:"type:varchar(255);not null;uniqueIndex;default:null"`
	Password string `gorm:"type:varchar(255);not null;default:null"`
}
