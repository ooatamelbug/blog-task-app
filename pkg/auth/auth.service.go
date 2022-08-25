package auth

import (
	"github.com/ooatamelbug/blog-task-app/pkg/auth/authdto"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
	"github.com/ooatamelbug/blog-task-app/pkg/users"
	"github.com/ooatamelbug/blog-task-app/pkg/users/dto"
)

type AuthService interface {
	SignUpUser(createUser dto.CreateUserDTO) (string, error)
	SignInUser(loginData authdto.Login) (string, error)
}

type authSerice struct {
	userService users.UserService
	jwtService  services.JWTService
}

func NewAuthService(userServe users.UserService, jwtServ services.JWTService) AuthService {
	return &authSerice{
		userService: userServe,
		jwtService:  jwtServ,
	}
}

func (authServ *authSerice) SignUpUser(createUser dto.CreateUserDTO) (string, error) {
	user, err := authServ.userService.CreateUser(createUser)
	var token string
	if err == nil {
		token = authServ.jwtService.GenerateToken(user.ID, user.Email)
	}
	return token, err
}

func (authServ *authSerice) SignInUser(loginData authdto.Login) (string, error) {
	user, err := authServ.userService.CredentialUser(loginData.Email, loginData.Password)
	var token string
	if err == nil {
		token = authServ.jwtService.GenerateToken(user.ID, user.Email)
		return token, nil
	}
	return "", err
}
