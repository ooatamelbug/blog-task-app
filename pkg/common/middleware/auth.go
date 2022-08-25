package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
)

const (
	AuthPayload = "payload_user"
)

func Auth(jwtService services.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			response := services.ReturnResponse(false, "not Authorization", nil, "", "no token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		authHeader = strings.Split(authHeader, " ")[1]
		payload, err := jwtService.ValidateToken(authHeader)
		if payload.Valid {
			claims := payload.Claims.(jwt.MapClaims)
			log.Println(claims)
		} else {
			log.Panicln(err)
			response := services.ReturnResponse(false, "not Authorization", nil, "", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		// ctx.Set(AuthPayload, payload)
		ctx.Next()
	}
}
