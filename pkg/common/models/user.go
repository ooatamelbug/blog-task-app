package models

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID           uint64     `gorm:"primary_key:auto_increment" json:"id"`
	UID          string     `gorm:"type: varchar(255)" json:"uid"`
	FirstName    string     `gorm:"type: varchar(255)" json:"firstname"`
	LastName     string     `gorm:"type: varchar(255)" json:"lastname"`
	Email        string     `gorm:"type: varchar(255)" json:"email"`
	Password     string     `gorm:"type: varchar(255);not null" json:"password"`
	PostsList    *[]Post    `json:"posts,omitempty"`
	CommentsList *[]Comment `json:"comments,omitempty"`
}
