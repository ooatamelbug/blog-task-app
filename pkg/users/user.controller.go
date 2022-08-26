package users

import (
	"github.com/gin-gonic/gin"
	"github.com/ooatamelbug/blog-task-app/pkg/common/middleware"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
)

type UserController interface {
	Index(ctx *gin.Context)
	GetProfile(ctx *gin.Context)
}

type userControllerData struct {
	userServ UserService
	jwtServ  services.JWTService
}

func NewUserController(userSer UserService, jwtserv services.JWTService) UserController {
	return &userControllerData{
		userServ: userSer,
		jwtServ:  jwtserv,
	}
}

func (userCont *userControllerData) Index(ctx *gin.Context) {
	resposne := services.ReturnResponse(true, "go go", "Hello in blog app (^-^)", "", "")
	ctx.JSON(200, resposne)
}

func (userCont *userControllerData) GetProfile(ctx *gin.Context) {
	userId := ctx.GetUint64(middleware.AuthPayload)
	user := userCont.userServ.ProfileUser(userId)
	resposne := services.ReturnResponse(true, "go", user, "", "")
	ctx.JSON(200, resposne)
}
