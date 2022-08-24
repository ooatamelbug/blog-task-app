package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ooatamelbug/blog-task-app/pkg/auth/authdto"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
	"github.com/ooatamelbug/blog-task-app/pkg/users/dto"
)

type AuthController interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type authController struct {
	authService AuthService
}

func NewAuthController(authuser AuthService) AuthController {
	return &authController{
		authService: authuser,
	}
}

func (authuser *authController) SignUp(ctx *gin.Context) {
	var createUser dto.CreateUserDTO
	errDto := ctx.ShouldBind(&createUser)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	token := authuser.authService.SignUpUser(createUser)
	if token == "" {
		response := services.ReturnResponse(false, "error in input data", nil, "", "error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", nil, token, "")
	ctx.JSON(http.StatusCreated, response)
}

func (authuser *authController) SignIn(ctx *gin.Context) {
	var loginUser authdto.Login
	errDto := ctx.ShouldBind(&loginUser)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	token := authuser.authService.SignInUser(loginUser)
	if token == "" {
		response := services.ReturnResponse(false, "error in input data", nil, "", "error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", nil, token, "")
	ctx.JSON(http.StatusCreated, response)
}
