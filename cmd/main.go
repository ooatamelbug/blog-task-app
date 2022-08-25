package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ooatamelbug/blog-task-app/pkg/auth"
	"github.com/ooatamelbug/blog-task-app/pkg/common/database"
	"github.com/ooatamelbug/blog-task-app/pkg/common/middleware"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
	"github.com/ooatamelbug/blog-task-app/pkg/posts"
	"github.com/ooatamelbug/blog-task-app/pkg/users"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB             = database.ConnectionDB()
	userRepository users.UserRepository = users.NewUserRepository(db)
	userService    users.UserService    = users.NewUserService(userRepository)
	userController users.UserController = users.NewUserController(userService)
	jwtService     services.JWTService  = services.NewJWTService()
	authService    auth.AuthService     = auth.NewAuthService(userService, jwtService)
	authController auth.AuthController  = auth.NewAuthController(authService)
	postRepository posts.PostRepository = posts.NewPostRepository(db)
	postService    posts.PostService    = posts.NewPostService(postRepository)
	postController posts.PostController = posts.NewPostController(postService, jwtService)
)

func main() {
	defer database.CloseConnectionDB(db)
	server := gin.Default()

	// user routes
	userRoutes := server.Group("/api/user")
	{
		userRoutes.GET("/", userController.Index)
	}

	// auth routes
	authRoutes := server.Group("/api/auth")
	{
		authRoutes.POST("/signin/", authController.SignIn)
		authRoutes.POST("/signup/", authController.SignUp)
	}

	// post routes
	postRoutes := server.Group("/api/post").Use(middleware.Auth(jwtService))
	{
		postRoutes.GET("/:id", postController.FindPost)
		postRoutes.GET("/all/", postController.GetAllPost)
		postRoutes.POST("/create/", postController.CreatePost)
		postRoutes.PUT("/update/", postController.UpdatePost)
		postRoutes.DELETE("/delete/", postController.DeletePost)
	}

	srv := &http.Server{
		Addr:    ":5000",
		Handler: server,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("faild to initilze server: %v\n", err)
		}
	}()
	log.Printf("listen on port %v\n", srv.Addr)
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("shutting down server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server force Shutdown %v\n", err)
	}
	// server.Run(":5000")
}
