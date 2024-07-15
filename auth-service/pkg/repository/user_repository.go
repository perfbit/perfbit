package repository

import (
	"database/sql"
	"github.com/maulikam/perfbit/auth-service/pkg/model"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	VerifyUser(username, code string) error
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, username, password, verified, code FROM users WHERE username = $1"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Verified, &user.Code)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) CreateUser(user *model.User) error {
	query := "INSERT INTO users (username, password, verified, code) VALUES ($1, $2, $3, $4)"
	_, err := r.db.Exec(query, user.Username, user.Password, user.Verified, user.Code)
	return err
}

func (r *PostgresUserRepository) VerifyUser(username, code string) error {
	query := "UPDATE users SET verified = true WHERE username = $1 AND code = $2"
	result, err := r.db.Exec(query, username, code)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
