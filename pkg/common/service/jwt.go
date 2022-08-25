package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userId uint64, email string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtUserClaim struct {
	UserID uint64 `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ydhrub",
		secretKey: os.Getenv("SECERTKEY"),
	}
}

func (jwtServ *jwtService) GenerateToken(UserID uint64, Email string) string {
	claims := &jwtUserClaim{
		UserID,
		Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    jwtServ.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(jwtServ.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (jwtServ *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected singin method %v\n", t.Header["alg"])
		}
		return []byte(jwtServ.secretKey), nil
	})
}
