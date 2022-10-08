package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:string;size:64;not null;comment:'用户姓名'"`
}
