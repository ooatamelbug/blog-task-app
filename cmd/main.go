package main

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
	db                *gorm.DB                   = database.ConnectionDB()
	userRepository    users.UserRepository       = users.NewUserRepository(db)
	userService       users.UserService          = users.NewUserService(userRepository)
	jwtService        services.JWTService        = services.NewJWTService()
	userController    users.UserController       = users.NewUserController(userService, jwtService)
	authService       auth.AuthService           = auth.NewAuthService(userService, jwtService)
	authController    auth.AuthController        = auth.NewAuthController(authService)
	postRepository    posts.PostRepository       = posts.NewPostRepository(db)
	postService       posts.PostService          = posts.NewPostService(postRepository)
	postController    posts.PostController       = posts.NewPostController(postService, jwtService)
	commentRepository comments.CommentRepository = comments.NewCommentRepository(db)
	commentService    comments.CommentService    = comments.NewCommentService(commentRepository)
	commentController comments.CommentController = comments.NewCommentController(commentService, jwtService)
)

func main() {
	defer database.CloseConnectionDB(db)
	server := gin.Default()

	// user routes
	userRoutes := server.Group("/api/user")
	{
		userRoutes.GET("/", userController.Index)
		userRoutes.GET("/data", userController.GetProfile)
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
		postRoutes.POST("/create", postController.CreatePost).Use(middleware.Auth(jwtService))
		postRoutes.PUT("/update/:id", postController.UpdatePost).Use(middleware.Auth(jwtService))
		postRoutes.DELETE("/delete/:id", postController.DeletePost).Use(middleware.Auth(jwtService))
	}

	// post routes
	commentRoutes := server.Group("/api/comment")
	{
		commentRoutes.GET("/one/:id", commentController.FindComment)
		commentRoutes.GET("/all", commentController.GetAllComment)
		commentRoutes.POST("/create", commentController.CreateComment).Use(middleware.Auth(jwtService))
		commentRoutes.PUT("/update/:id", commentController.UpdateComment).Use(middleware.Auth(jwtService))
		commentRoutes.DELETE("/delete/:id", commentController.DeleteComment).Use(middleware.Auth(jwtService))
	}

	server.Run(":5000")
}
