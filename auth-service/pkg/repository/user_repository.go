// pkg/repository/user_repository.go
package repository

import (
	"database/sql"
	"github.com/maulikam/auth-service/pkg/model"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) FindByUsername(username string) (*model.User, error) {
	user := &model.User{}
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) CreateUser(user *model.User) error {
	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err := r.db.Exec(query, user.Username, user.Password)
	return err
}
