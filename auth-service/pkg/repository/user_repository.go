package repository

import (
	"database/sql"
	"github.com/maulikam/perfbit/auth-service/pkg/model"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	VerifyUser(username, code string) error
	UpdateRefreshToken(username, refreshToken string) error
	GetUserByGitHubUsername(gitHubUsername string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, username, password, verified, code, refresh_token FROM users WHERE username = $1"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password, &user.Verified, &user.Code, &user.RefreshToken)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) CreateUser(user *model.User) error {
	query := "INSERT INTO users (username, password, github_username, verified, code, refresh_token) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := r.db.Exec(query, user.Username, user.Password, user.GitHubUsername, user.Verified, user.Code, user.RefreshToken)
	return err
}

func (r *PostgresUserRepository) UpdateUser(user *model.User) error {
	query := "UPDATE users SET username = $1, password = $2, github_username = $3, verified = $4, code = $5, refresh_token = $6 WHERE id = $7"
	_, err := r.db.Exec(query, user.Username, user.Password, user.GitHubUsername, user.Verified, user.Code, user.RefreshToken, user.ID)
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

func (r *PostgresUserRepository) UpdateRefreshToken(username, refreshToken string) error {
	query := "UPDATE users SET refresh_token = $1 WHERE username = $2"
	_, err := r.db.Exec(query, refreshToken, username)
	return err
}

// repository/user_repository.go
func (r *PostgresUserRepository) GetUserByGitHubUsername(gitHubUsername string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, username, password, github_username, verified, code, refresh_token FROM users WHERE github_username = $1"
	err := r.db.QueryRow(query, gitHubUsername).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.GitHubUsername,
		&user.Verified,
		&user.Code,
		&user.RefreshToken,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) GetUserByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, username, password, github_username, verified, code, refresh_token FROM users WHERE username = $1"
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.GitHubUsername,
		&user.Verified,
		&user.Code,
		&user.RefreshToken,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
