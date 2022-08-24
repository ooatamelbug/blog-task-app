package users

import (
	"log"

	"github.com/ooatamelbug/blog-task-app/pkg/users/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user dto.UserCreateDTO) User
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) CreateUser(user dto.UserCreateDTO) User {

}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("faild to hash")
	}
	return string(hash)
}

func ComparePassword(hashedPassword string, plantextPassword []byte) bool {
	bytehash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(bytehash, plantextPassword)

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
