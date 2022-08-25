package users

import (
	"github.com/gin-gonic/gin"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
)

type UserController interface {
	Index(ctx *gin.Context)
}

type userControllerData struct {
	userServ UserService
}

func NewUserController(userSer UserService) UserController {
	return &userControllerData{
		userServ: userSer,
	}
}

func (userCont *userControllerData) Index(ctx *gin.Context) {
	resposne := services.ReturnResponse(true, "go", "here", "", "")
	ctx.JSON(200, resposne)
}
