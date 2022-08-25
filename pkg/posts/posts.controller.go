package posts

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
	"github.com/ooatamelbug/blog-task-app/pkg/posts/dto"
)

type PostController interface {
	CreatePost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
	FindPost(ctx *gin.Context)
	GetAllPost(ctx *gin.Context)
	GetUserIdByToken(token string) (uint64, error)
}

type postController struct {
	posyservice PostService
	jwtServ     services.JWTService
}

func NewPostController(postServ PostService, jwtservice services.JWTService) PostController {
	return &postController{
		posyservice: postServ,
		jwtServ:     jwtservice,
	}
}

func (postControl *postController) CreatePost(ctx *gin.Context) {
	var createPost dto.CreatePostDto
	errDto := ctx.BindJSON(&createPost)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	userId, err := postControl.GetUserIdByToken(ctx.GetHeader("Authorization"))
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	createPost.User = userId
	post, err := postControl.posyservice.CreatePost(createPost)
	if post.Title == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (postControl *postController) UpdatePost(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 0, 0)
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	var createPost dto.CreatePostDto
	errDto := ctx.BindJSON(&createPost)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	userId, errToken := postControl.GetUserIdByToken(ctx.GetHeader("Authorization"))
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errToken.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	createPost.User = userId
	post, err := postControl.posyservice.UpdatePost(createPost, idUint)
	if post.Title == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (postControl *postController) DeletePost(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 0, 0)
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	userId, err := postControl.GetUserIdByToken(ctx.GetHeader("Authorization"))
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	post, err := postControl.posyservice.DeletePost(idUint, userId)
	if post.Title == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (postControl *postController) FindPost(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	post, err := postControl.posyservice.GetPost(idUint)
	if post.Title == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (postControl *postController) GetAllPost(ctx *gin.Context) {
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 0, 0)
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	post, err := postControl.posyservice.GetPost(idUint)
	if post.Title == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (postControl *postController) GetUserIdByToken(token string) (uint64, error) {
	var userId uint64
	payload, err := postControl.jwtServ.ValidateToken(token)
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
