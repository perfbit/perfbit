// pkg/repository/user_repository.go
package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/maulikam/auth-service/pkg/model"
)

type UserRepository interface {
	FindByUsername(username string) (*model.User, error)
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
