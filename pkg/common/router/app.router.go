package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ooatamelbug/blog-task-app/pkg/auth"
	"github.com/ooatamelbug/blog-task-app/pkg/comments"
	"github.com/ooatamelbug/blog-task-app/pkg/common/database"
	"github.com/ooatamelbug/blog-task-app/pkg/common/middleware"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
	"github.com/ooatamelbug/blog-task-app/pkg/posts"
	"github.com/ooatamelbug/blog-task-app/pkg/users"
	"gorm.io/gorm"
)

var (
	DB                *gorm.DB                   = database.ConnectionDB()
	userRepository    users.UserRepository       = users.NewUserRepository(DB)
	userService       users.UserService          = users.NewUserService(userRepository)
	jwtService        services.JWTService        = services.NewJWTService()
	userController    users.UserController       = users.NewUserController(userService, jwtService)
	authService       auth.AuthService           = auth.NewAuthService(userService, jwtService)
	authController    auth.AuthController        = auth.NewAuthController(authService)
	postRepository    posts.PostRepository       = posts.NewPostRepository(DB)
	postService       posts.PostService          = posts.NewPostService(postRepository)
	postController    posts.PostController       = posts.NewPostController(postService, jwtService)
	commentRepository comments.CommentRepository = comments.NewCommentRepository(DB)
	commentService    comments.CommentService    = comments.NewCommentService(commentRepository)
	commentController comments.CommentController = comments.NewCommentController(commentService, jwtService)
)

func HandelRouter(server *gin.Engine) {

	// index
	server.GET("api/", userController.Index)

	// user routes
	userRoutes := server.Group("/api/user")
	{
		userRoutes.GET("/data", middleware.Auth(jwtService), userController.GetProfile)
	}

	// auth routes
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/signin/", authController.SignIn)
		authRoutes.POST("/signup/", authController.SignUp)
	}

	// post routes
	postRoutes := server.Group("/api/post")
	{
		postRoutes.GET("/one/:id", postController.FindPost)
		postRoutes.GET("/all", postController.GetAllPost)
		postRoutes.POST("/create", middleware.Auth(jwtService), postController.CreatePost)
		postRoutes.PUT("/update/:id", middleware.Auth(jwtService), postController.UpdatePost)
		postRoutes.DELETE("/delete/:id", middleware.Auth(jwtService), postController.DeletePost)
	}

	// post routes
	commentRoutes := server.Group("/api/comment")
	{
		commentRoutes.GET("/one/:id", commentController.FindComment)
		commentRoutes.GET("/all", middleware.Auth(jwtService), commentController.GetAllComment)
		commentRoutes.POST("/create", middleware.Auth(jwtService), commentController.CreateComment)
		commentRoutes.PUT("/update/:id", middleware.Auth(jwtService), commentController.UpdateComment)
		commentRoutes.DELETE("/delete/:id", middleware.Auth(jwtService), commentController.DeleteComment)
	}

}
