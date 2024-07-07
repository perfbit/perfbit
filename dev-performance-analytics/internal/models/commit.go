package models

import (
    "gorm.io/gorm"
)

type Commit struct {
    gorm.Model
    BranchID uint   `gorm:"not null"`
    Hash     string `gorm:"uniqueIndex;not null"`
    Message  string `gorm:"not null"`
    Author   string `gorm:"not null"`
}
