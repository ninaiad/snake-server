package database

import (
	"fmt"
	"snake"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (auth_postgres *AuthPostgres) CreateUser(user snake.User) (error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, score) values ($1, $2, $3) RETURNING id", usersTable)

	row := auth_postgres.db.QueryRow(query, user.Username, user.Password, user.Score)
	if err := row.Scan(&id); err != nil {
		return err
	}

	return nil
}

func (auth_postgres *AuthPostgres) GetUser(username, password string) (snake.User, error) {
	var user snake.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := auth_postgres.db.Get(&user, query, username, password)

	return user, err
}

func (auth_postgres *AuthPostgres) DeleteUser(userId int) (error) {
	query := fmt.Sprintf("DELETE FROM %s ul WHERE ul.id=$1",
		usersTable)
	_, err := auth_postgres.db.Exec(query, userId)

	return err
}
