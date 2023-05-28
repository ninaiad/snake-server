package service

import (
	"snake"
	"snake/pkg/database"
)

type ScoreService struct {
	db database.Score
}

func NewScoreService(db database.Score) *ScoreService {
	return &ScoreService{db: db}
}

func (s *ScoreService) GetAllScores() ([]snake.UserPublic, error) {
	return s.db.GetAllScores()
}

func (s *ScoreService) GetScore(userId int) (snake.UserPublic, error) { // a json similar to snake.User but without passwords!!!
	return s.db.GetScore(userId)
}

func (s *ScoreService) Update(userId int, value uint64) error {
	return s.db.Update(userId, value)
}
