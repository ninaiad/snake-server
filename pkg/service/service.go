package service

import (
	"snake"
	"snake/pkg/database"
)

type Authorization interface {
	CreateUser(user snake.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	DeleteUser(userId int) error
}

type Score interface {
	GetAllScores() ([]snake.UserPublic, error)
	GetScore(userId int) (snake.UserPublic, error)
	Update(userId int, value uint64) error
}

type Service struct {
	Authorization
	Score
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization: CreateAuthService(db.Authorization),
		Score:  NewScoreService(db.Score),
	}
}
