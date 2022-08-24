package users

import "time"

type User struct {
	ID        uint64 `gorm:"primary_key:auto_increment" json:"id"`
	UID       string `gorm:"type: varchar(255)" json:"uid"`
	FirstName string `gorm:"type: varchar(255)" json:"firstname"`
	LastName  string `gorm:"type: varchar(255)" json:"lastname"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Password  string `gorm:"->: <-;not null" json:"password"`
	CreateAt  time.Time
	UpdateAt  time.Time
}
