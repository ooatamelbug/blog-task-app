package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID    uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Body  string `gorm:"type: text" json:"body"`
	Likes int64  `gorm:"type: text" json:"likes"`
	User  User   `gorm:"foreignkey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	Post  Post   `gorm:"foreignkey:PostID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"post"`
}
