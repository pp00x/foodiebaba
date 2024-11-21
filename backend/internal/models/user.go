package models

import (
    "gorm.io/gorm"
    "time"
)

type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    Username  string         `gorm:"uniqueIndex;size:100" json:"username" validate:"required"`
    Email     string         `gorm:"uniqueIndex;size:100" json:"email" validate:"required,email"`
    Password  string         `gorm:"size:255" json:"password" validate:"required"`
    Role      string         `gorm:"size:20" json:"role"` // "user" or "admin"
    Reputation int           `json:"reputation" gorm:"default:0"`
    Listings  []Restaurant   `json:"listings" gorm:"foreignKey:CreatedByID"`
    Reviews   []Review       `json:"reviews"`
}