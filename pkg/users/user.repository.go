package users

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user User) User
	FindOne(userId string) User
	FindAll() []User
	Update(user User) User
	Delete(user User) User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Create(user User) User {
	db.connection.Save(user)
	return user
}

func (db *userConnection) Update(user User) User {
	db.connection.Save(user)
	return user
}

func (db *userConnection) Delete(user User) User {
	db.connection.Delete(user)
	return user
}

func (db *userConnection) FindOne(userId string) User {
	var user User
	db.connection.Where("id = ?", userId).Take(&user)
	return user
}

func (db *userConnection) FindAll() []User {
	var user []User
	db.connection.Find(&user)
	return user
}
