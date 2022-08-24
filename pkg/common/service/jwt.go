package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JwtService interface {
	GenerateToken(userId uint64, email string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtUserClaim struct {
	UserID uint64 `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type jwtService struct {
	secertKey string
	issuer    string
}

func NewJwtService(typeJwt string) JwtService {
	return &jwtService{
		issuer:    "ydhrub",
		secertKey: os.Getenv("SECERTKEY"),
	}
}

func (jwtServ *jwtService) GenerateToken(userId uint64, email string) string {
	claims := &jwtUserClaim{
		UserID: userId,
		Email:  email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    jwtServ.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	t, err := token.SignedString([]byte(jwtServ.secertKey))
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
		return []byte(jwtServ.secertKey), nil
	})
}
