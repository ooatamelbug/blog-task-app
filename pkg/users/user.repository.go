package users

import (
	"github.com/ooatamelbug/blog-task-app/pkg/common/models"
	"github.com/ooatamelbug/blog-task-app/pkg/users/dto"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindOne(search dto.SearchUser) models.User
	FindAll() []models.User
	Update(user models.User) (models.User, error)
	Delete(user models.User) models.User
	FindAnd(searchWithAnd dto.SearchWithAnd) models.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Create(user models.User) (models.User, error) {
	err := db.connection.Save(&user)
	return user, err.Error
}

func (db *userConnection) Update(user models.User) (models.User, error) {
	err := db.connection.Save(&user)
	return user, err.Error
}

func (db *userConnection) Delete(user models.User) models.User {
	db.connection.Delete(user)
	return user
}

func (db *userConnection) FindOne(search dto.SearchUser) models.User {
	var user models.User
	db.connection.Preload("Post").Preload("CommentsList").Where("email = ?", search.Email).Find(&user)
	return user
}

func (db *userConnection) FindAll() []models.User {
	var user []models.User
	db.connection.Preload("PostsList").Preload("CommentsList").Find(&user)
	return user
}

func (db *userConnection) FindAnd(searchWithAnd dto.SearchWithAnd) models.User {
	var user models.User
	db.connection.Preload("PostsList").Preload("CommentsList").Where(&models.User{ID: searchWithAnd.ID}).First(&user)
	return user
}
