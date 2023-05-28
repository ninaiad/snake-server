package database

import (
	"fmt"
	"snake"

	"github.com/jmoiron/sqlx"
)

type ScorePostgres struct {
	db *sqlx.DB
}

func NewScorePostgres(db *sqlx.DB) *ScorePostgres {
	return &ScorePostgres{db: db}
}

func (r *ScorePostgres) GetAllScores() ([]snake.UserPublic, error) {
	var scores []snake.UserPublic

	query := fmt.Sprintf("SELECT ul.username, ul.score FROM %s ul", usersTable)
	err := r.db.Select(&scores, query)

	return scores, err
}

func (r *ScorePostgres) GetScore(userId int) (snake.UserPublic, error) {
	var score snake.UserPublic

	query := fmt.Sprintf("SELECT ul.username, ul.score FROM %s ul WHERE ul.id = $1",
		usersTable)
	err := r.db.Get(&score, query, userId)

	return score, err
}

func (r *ScorePostgres) Update(userId int, value uint64) error {
	var score uint64
	query := fmt.Sprintf("SELECT ul.score FROM %s ul WHERE ul.id = $1", usersTable)
	err := r.db.Get(&score, query, userId)
	if err != nil {
		return err
	}

	score += value

	query = fmt.Sprintf("UPDATE %s ul SET score = $1 WHERE ul.id = $2", usersTable)
	_, err = r.db.Exec(query, score, userId)

	return err
}
