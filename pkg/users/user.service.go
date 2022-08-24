package users

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/ooatamelbug/blog-task-app/pkg/users/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user dto.CreateUserDTO) User
	GetOneUserByEmail(email string) (User, error)
	UpdateUser(user dto.UpdateUserDTO) User
	CredentialUser(email string, password string) User
	ProfileUser(userId uint64) User
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) GetOneUserByEmail(email string) (User, error) {
	search := dto.SearchUser{}
	search.Email = email
	row, result := service.userRepository.FindOne(search)
	// if result.Error != nil {
	// 	log.Fatalf("error in get data %v\n", result.Error)
	// }
	if result.Error != nil {
		return row, result.Error
	}
	return row, nil
}

func (service *userService) CreateUser(user dto.CreateUserDTO) User {
	newUser := User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("error in map %v\n", err)
	}

	row, err := service.GetOneUserByEmail(user.Email)
	if err == nil {
		fmt.Printf("%+v\n", row)
		log.Fatalf("this email is used %v\n", err)
	}

	hashPassword := hashAndSalt([]byte(newUser.Password))
	newUser.Password = hashPassword

	res := service.userRepository.Create(newUser)
	return res
}

func (service *userService) UpdateUser(user dto.UpdateUserDTO) User {
	userToUpdate := User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("error in map %v\n", err)
	}

	if user.Email != "" {
		row, err := service.GetOneUserByEmail(user.Email)
		if err == nil && row.ID != user.ID {
			fmt.Printf("%+v\n", row)
			log.Fatalf("this email is used %v\n", err)
		}
	}

	if user.Password != "" {
		hashPassword := hashAndSalt([]byte(user.Password))
		userToUpdate.Password = hashPassword
	}

	service.userRepository.Update(userToUpdate)
	return userToUpdate
}

func (sercive *userService) CredentialUser(email string, password string) User {
	userSearchData := dto.SearchUser{}
	userSearchData.Email = email
	userCred, res := sercive.userRepository.FindOne(userSearchData)
	if res.Error != nil {
		panic("error in email")
	}

	if comparePasswordFor := ComparePassword(userCred.Password, []byte(password)); !comparePasswordFor {
		panic("password error")
	}
	return userCred
}

func (sercive *userService) ProfileUser(userId uint64) User {
	getById := dto.SearchWithAnd{}
	getById.ID = userId
	row := sercive.userRepository.FindAnd(getById)
	return row
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
