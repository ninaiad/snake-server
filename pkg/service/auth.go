package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"snake"
	"snake/pkg/database"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var salt = os.Getenv("PASSWORD_SALT")
var signingKey = os.Getenv("TOKEN_SIGNING_KEY")

const (
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	db database.Authorization
}

func CreateAuthService(db database.Authorization) *AuthService {
	return &AuthService{db: db}
}

func (auth_service *AuthService) CreateUser(user snake.User) error {
	user.Password = generatePasswordHash(user.Password)
	user.Score = 0
	return auth_service.db.CreateUser(user)
}

func (auth_service *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := auth_service.db.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (auth_service *AuthService) DeleteUser(userId int) error {
	return auth_service.db.DeleteUser(userId)
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
