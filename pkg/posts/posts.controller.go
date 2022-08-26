package posts

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ooatamelbug/blog-task-app/pkg/common/helper"
	"github.com/ooatamelbug/blog-task-app/pkg/common/middleware"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
	"github.com/ooatamelbug/blog-task-app/pkg/posts/dto"
)

type PostController interface {
	CreatePost(ctx *gin.Context)
	UpdatePost(ctx *gin.Context)
	DeletePost(ctx *gin.Context)
	FindPost(ctx *gin.Context)
	GetAllPost(ctx *gin.Context)
}

type postController struct {
	postService PostService
	jwtServ     services.JWTService
}

func NewPostController(postServ PostService, jwtservice services.JWTService) PostController {
	return &postController{
		postService: postServ,
		jwtServ:     jwtservice,
	}
}

func (postControl *postController) CreatePost(ctx *gin.Context) {
	userId := ctx.GetUint64(middleware.AuthPayload)
	var createPost dto.CreatePostDto

	createPost.UserID = userId
	errDto := ctx.BindJSON(&createPost)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	post, err := postControl.postService.CreatePost(createPost)
	if post.Title == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (postControl *postController) UpdatePost(ctx *gin.Context) {
	idUint, err := helper.ConvertToInt(ctx.Param("id"))
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var createPost dto.CreatePostDto
	userId := ctx.GetUint64(middleware.AuthPayload)

	createPost.UserID = userId

	errDto := ctx.BindJSON(&createPost)
	if errDto != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", errDto.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	post, err := postControl.postService.UpdatePost(createPost, idUint)
	if post.Title == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (postControl *postController) DeletePost(ctx *gin.Context) {
	idUint, err := helper.ConvertToInt(ctx.Param("id"))
	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	userId := ctx.GetUint64(middleware.AuthPayload)

	post, err := postControl.postService.DeletePost(idUint, userId)
	if post.Title == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (postControl *postController) FindPost(ctx *gin.Context) {
	idUint, err := helper.ConvertToInt(ctx.Param("id"))

	if err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}
	post, err := postControl.postService.GetPost(idUint)
	if post.Title == "" && err != nil {
		response := services.ReturnResponse(false, "error in input data", nil, "", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}

func (postControl *postController) GetAllPost(ctx *gin.Context) {
	post := postControl.postService.GetPosts()
	if length := len(post); length < 0 {
		response := services.ReturnResponse(false, "error in input data", nil, "", "error")
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := services.ReturnResponse(true, "go", post, "", "")
	ctx.JSON(http.StatusCreated, response)
}
