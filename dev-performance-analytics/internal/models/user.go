package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex;not null"`
    Password string `gorm:"not null"`
    Email    string `gorm:"uniqueIndex;not null"`
}
