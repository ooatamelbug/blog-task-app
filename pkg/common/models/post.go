package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID    uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UID   string `gorm:"type: varchar(255)" json:"uid"`
	Title string `gorm:"type: text" json:"title"`
	Body  string `gorm:"type: text" json:"body"`
	Likes int64  `gorm:"type: text" json:"likes"`
	User  User   `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
}
