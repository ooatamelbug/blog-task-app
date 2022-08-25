package models

import "gorm.io/gorm"

type Post struct {
	*gorm.Model
	ID           uint64     `gorm:"primary_key:auto_increment" json:"id"`
	Title        string     `gorm:"type:text" json:"title"`
	Body         string     `gorm:"type:text" json:"body"`
	Likes        int64      `gorm:"type:int; default:0" json:"likes"`
	UserID       uint64     `gorm:"not null" json:"-"`
	User         User       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	CommentsList *[]Comment `json:"comments,omitempty"`
}
