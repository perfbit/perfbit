package models

import (
    "gorm.io/gorm"
)

type Branch struct {
    gorm.Model
    RepositoryID uint   `gorm:"not null"`
    Name         string `gorm:"not null"`
    Commits      []Commit
}
