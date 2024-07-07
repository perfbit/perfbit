package models

import (
    "gorm.io/gorm"
)

type Repository struct {
    gorm.Model
    Owner   string `gorm:"not null"`
    Name    string `gorm:"not null"`
    Branches []Branch
}
