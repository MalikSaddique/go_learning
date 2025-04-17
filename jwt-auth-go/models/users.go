package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int      `json:"id"`
	Email    string   `json:"email" gorm:"unique"`
	Password string   `json:"password"`
	Results  []Result `gorm:"foreignKey:UserID"`
}
