package models

import "gorm.io/gorm"

type Comment struct {
	*gorm.Model
	ID     uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Body   string `gorm:"type:text" json:"body"`
	Likes  int64  `gorm:"type:int; default:0" json:"likes"`
	UserID uint64 `gorm:"not null" json:"-"`
	User   User   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	PostID uint64 `gorm:"not null" json:"-"`
	Post   Post   `gorm:"foreignKey:PostID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"post"`
}
