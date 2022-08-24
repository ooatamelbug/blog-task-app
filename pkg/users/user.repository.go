package users

import (
	"github.com/ooatamelbug/blog-task-app/pkg/users/dto"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user User) User
	FindOne(search dto.SearchUser) (User, *gorm.DB)
	FindAll() []User
	Update(user User) User
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

func (db *userConnection) FindOne(search dto.SearchUser) (User, *gorm.DB) {
	var user User
	result := db.connection.Where("id = ?", search.UserId).Or("email = ?", search.Email).Find(&user)
	return user, result
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
