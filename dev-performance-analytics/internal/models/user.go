package models

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex;not null"`
    Password string `gorm:"not null"`
    Email    string `gorm:"uniqueIndex;not null"`
    GithubID int64  `gorm:"uniqueIndex;not null"` // Added GithubID
    Login    string `gorm:"not null"`             // Added Login
    Name     string                              // Added Name
}
