package database

import (
	"snake"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user snake.User) (error)
	GetUser(username, password string) (snake.User, error)
	DeleteUser(userId int) (error)
}

type Score interface {
	GetAllScores() ([]snake.UserPublic, error)
	GetScore(userId int) (snake.UserPublic, error)
	Update(userId int, value uint64) error
}

type Database struct {
	Authorization
	Score
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Authorization: NewAuthPostgres(db),
		Score: NewScorePostgres(db),
	}
}
