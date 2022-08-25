package users

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	resposne := services.ReturnResponse(true, "go", "here", "", "")
	ctx.JSON(200, resposne)
}

func (userCont *userControllerData) GetProfile(ctx *gin.Context) {
	userId, errToken := userCont.GetUserIdByToken(strings.Split(ctx.GetHeader("Authorization"), " ")[1])
	if errToken != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errToken.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	user := userCont.userServ.ProfileUser(userId)
	resposne := services.ReturnResponse(true, "go", user, "", "")
	ctx.JSON(200, resposne)
}

func (userCont *userControllerData) GetUserIdByToken(token string) (uint64, error) {
	var userId uint64
	payload, err := userCont.jwtServ.ValidateToken(token)
	if err != nil {
		return userId, err
	}
	claims := payload.Claims.(jwt.MapClaims)
	d := fmt.Sprintf("%v", claims["user_id"])
	idUint, err := strconv.ParseUint(d, 0, 0)
	if err != nil {
		return userId, err
	}
	return idUint, nil
}
