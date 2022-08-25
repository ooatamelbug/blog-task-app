package users

import (
	"errors"
	"log"
	"os/exec"

	"github.com/mashingan/smapping"
	"github.com/ooatamelbug/blog-task-app/pkg/users/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user dto.CreateUserDTO) (User, error)
	GetOneUserByEmail(email string) User
	UpdateUser(user dto.UpdateUserDTO) (User, error)
	CredentialUser(email string, password string) (User, error)
	ProfileUser(userId uint64) User
}

type userServiceData struct {
	userRepository UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userServiceData{
		userRepository: userRepo,
	}
}

func (service *userServiceData) GetOneUserByEmail(email string) User {
	search := dto.SearchUser{}
	search.Email = email
	row := service.userRepository.FindOne(search)
	return row
}

func (service *userServiceData) CreateUser(user dto.CreateUserDTO) (User, error) {
	newUser := User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("error in map %v\n", err)
		return newUser, err
	}

	row := service.GetOneUserByEmail(user.Email)
	if row.Email != "" {
		return row, errors.New("this email is used before")
	}

	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		return newUser, err
	}
	newUser.UID = string(out)

	hashPassword := hashAndSalt([]byte(newUser.Password))
	newUser.Password = hashPassword

	return service.userRepository.Create(newUser)

}

func (service *userServiceData) UpdateUser(user dto.UpdateUserDTO) (User, error) {
	userToUpdate := User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("error in map %v\n", err)
	}

	if user.Email != "" {
		row := service.GetOneUserByEmail(user.Email)
		if row.ID != user.ID {
			return userToUpdate, errors.New("this email is used before")
		}
	}

	if user.Password != "" {
		hashPassword := hashAndSalt([]byte(user.Password))
		userToUpdate.Password = hashPassword
	}

	return service.userRepository.Update(userToUpdate)
}

func (sercive *userServiceData) CredentialUser(email string, password string) (User, error) {
	userSearchData := dto.SearchUser{}
	userSearchData.Email = email
	userCred := sercive.userRepository.FindOne(userSearchData)
	if userCred.Email == "" {
		return userCred, errors.New("this email is not correct")
	}

	if comparePasswordFor := ComparePassword(userCred.Password, []byte(password)); !comparePasswordFor {
		return userCred, errors.New("this password is not correct")
	}
	return userCred, nil
}

func (sercive *userServiceData) ProfileUser(userId uint64) User {
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
