package repository

import (
    "dev-performance-analytics/internal/models"
    "gorm.io/gorm"
)

type RepositoryRepository interface {
    CreateRepository(repo *models.Repository) error
    GetRepositoryByID(id uint) (*models.Repository, error)
}

type repositoryRepository struct {
    db *gorm.DB
}

func NewRepositoryRepository(db *gorm.DB) RepositoryRepository {
    return &repositoryRepository{db: db}
}

func (r *repositoryRepository) CreateRepository(repo *models.Repository) error {
    return r.db.Create(repo).Error
}

func (r *repositoryRepository) GetRepositoryByID(id uint) (*models.Repository, error) {
    var repo models.Repository
    err := r.db.Preload("Branches.Commits").First(&repo, id).Error
    return &repo, err
}
