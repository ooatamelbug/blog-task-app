package users

import (
	"github.com/ooatamelbug/blog-task-app/pkg/users/dto"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user User) (User, error)
	FindOne(search dto.SearchUser) User
	FindAll() []User
	Update(user User) (User, error)
	Delete(user User) User
	FindAnd(searchWithAnd dto.SearchWithAnd) User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Create(user User) (User, error) {
	err := db.connection.Save(&user)
	return user, err.Error
}

func (db *userConnection) Update(user User) (User, error) {
	err := db.connection.Save(&user)
	return user, err.Error
}

func (db *userConnection) Delete(user User) User {
	db.connection.Delete(user)
	return user
}

func (db *userConnection) FindOne(search dto.SearchUser) User {
	var user User
	db.connection.Where("email = ?", search.Email).Find(&user)
	return user
}

func (db *userConnection) FindAll() []User {
	var user []User
	db.connection.Find(&user)
	return user
}

func (db *userConnection) FindAnd(searchWithAnd dto.SearchWithAnd) User {
	var user User
	db.connection.Where(&User{ID: searchWithAnd.ID}).First(&user)
	return user
}
