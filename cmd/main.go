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
	"github.com/ooatamelbug/blog-task-app/pkg/common/database"
	"github.com/ooatamelbug/blog-task-app/pkg/common/router"
)

func main() {
	defer database.CloseConnectionDB(router.DB)
	server := gin.Default()

	// handel all router
	router.HandelRouter(server)

	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
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
}
