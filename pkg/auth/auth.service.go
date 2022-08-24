package auth

import (
	"github.com/ooatamelbug/blog-task-app/pkg/auth/authdto"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
	"github.com/ooatamelbug/blog-task-app/pkg/users"
	"github.com/ooatamelbug/blog-task-app/pkg/users/dto"
)

type AuthService interface {
	SignUpUser(createUser dto.CreateUserDTO) string
	SignInUser(loginData authdto.Login) string
}

type authSerice struct {
	userService users.UserService
	jwtService  services.JwtService
}

func NewAuthService(userServe users.UserService, jwtServ services.JwtService) AuthService {
	return &authSerice{
		userService: userServe,
		jwtService:  jwtServ,
	}
}

func (authServ *authSerice) SignUpUser(createUser dto.CreateUserDTO) string {
	user := authServ.userService.CreateUser(createUser)
	token := authServ.jwtService.GenerateToken(user.ID, user.Email)
	return token
}

func (authServ *authSerice) SignInUser(loginData authdto.Login) string {
	user := authServ.userService.CredentialUser(loginData.Email, loginData.Password)
	token := authServ.jwtService.GenerateToken(user.ID, user.Email)
	return token
}
