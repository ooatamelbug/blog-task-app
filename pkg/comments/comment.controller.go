package comments

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ooatamelbug/blog-task-app/pkg/comments/dto"
	"github.com/ooatamelbug/blog-task-app/pkg/common/helper"
	"github.com/ooatamelbug/blog-task-app/pkg/common/middleware"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
	FindComment(ctx *gin.Context)
	GetAllComment(ctx *gin.Context)
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
	userId := ctx.GetUint64(middleware.AuthPayload)

	createComment.UserID = userId

	errDto := ctx.BindJSON(&createComment)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	comment, err := commentControl.posyservice.CreateComment(createComment)
	if comment.Body == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (commentControl *commentController) UpdateComment(ctx *gin.Context) {
	idUint, err := helper.ConvertToInt(ctx.Param("id"))

	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var createComment dto.CreateCommentDto
	userId := ctx.GetUint64(middleware.AuthPayload)

	createComment.UserID = userId

	errDto := ctx.BindJSON(&createComment)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	comment, err := commentControl.posyservice.UpdateComment(createComment, idUint)
	if comment.Body == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (commentControl *commentController) DeleteComment(ctx *gin.Context) {
	idUint, err := helper.ConvertToInt(ctx.Param("id"))

	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := ctx.GetUint64(middleware.AuthPayload)

	comment, err := commentControl.posyservice.DeleteComment(idUint, userId)
	if comment.Body == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (commentControl *commentController) FindComment(ctx *gin.Context) {
	idUint, err := helper.ConvertToInt(ctx.Param("id"))

	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	comment, err := commentControl.posyservice.GetComment(idUint)
	if comment.Body == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (commentControl *commentController) GetAllComment(ctx *gin.Context) {
	comment := commentControl.posyservice.GetComments()
	if length := len(comment); length < 0 {
		response := services.ReturnResponse(false, "error in input data", nil, "", "error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := services.ReturnResponse(true, "go", comment, "", "")
	ctx.JSON(http.StatusCreated, response)
}
