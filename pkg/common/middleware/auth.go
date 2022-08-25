package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	services "github.com/ooatamelbug/blog-task-app/pkg/common/service"
)

func Auth(jwtService services.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")[1]

		if authHeader == "" {
			response := services.ReturnResponse(false, "not Authorization", nil, "", "no token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Panicln("claims[user_id]:", claims["user_id"])
			log.Panicln("claims[issuer]:", claims["issuer"])
		} else {
			log.Panicln(err)
			response := services.ReturnResponse(false, "not Authorization", nil, "", err.Error())
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		ctx.Next()
	}
}
