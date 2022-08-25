package comments

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ooatamelbug/blog-task-app/pkg/comments/dto"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
	FindComment(ctx *gin.Context)
	GetAllComment(ctx *gin.Context)
	GetUserIdByToken(token string) (uint64, error)
}

type commentController struct {
	posyservice CommentService
	jwtServ     services.JWTService
}

func NewCommentController(commentServ CommentService, jwtservice services.JWTService) CommentController {
	return &commentController{
		posyservice: commentServ,
		jwtServ:     jwtservice,
	}
}

func (commentControl *commentController) CreateComment(ctx *gin.Context) {
	var createComment dto.CreateCommentDto
	errDto := ctx.BindJSON(&createComment)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	userId, err := commentControl.GetUserIdByToken(ctx.GetHeader("Authorization"))
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	createComment.User = userId
	comment, err := commentControl.posyservice.CreateComment(createComment)
	if comment.Body == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (commentControl *commentController) UpdateComment(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 0, 0)
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	var createComment dto.CreateCommentDto
	errDto := ctx.BindJSON(&createComment)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	userId, errToken := commentControl.GetUserIdByToken(ctx.GetHeader("Authorization"))
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errToken.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	createComment.User = userId
	comment, err := commentControl.posyservice.UpdateComment(createComment, idUint)
	if comment.Body == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (commentControl *commentController) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 0, 0)
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	userId, err := commentControl.GetUserIdByToken(ctx.GetHeader("Authorization"))
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	comment, err := commentControl.posyservice.DeleteComment(idUint, userId)
	if comment.Body == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (commentControl *commentController) FindComment(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	comment, err := commentControl.posyservice.GetComment(idUint)
	if comment.Body == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (commentControl *commentController) GetAllComment(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 0, 0)
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	comment, err := commentControl.posyservice.GetComment(idUint)
	if comment.Body == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (commentControl *commentController) GetUserIdByToken(token string) (uint64, error) {
	var userId uint64
	payload, err := commentControl.jwtServ.ValidateToken(token)
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
