package models

import (
    "gorm.io/gorm"
    "time"
)

type Restaurant struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    Name        string         `gorm:"size:255" json:"name" validate:"required"`
    Address     string         `gorm:"size:255" json:"address" validate:"required"`
    Category    string         `gorm:"size:100" json:"category" validate:"required"`
    Description string         `gorm:"type:text" json:"description" validate:"required"`
    Photos      []Photo        `json:"photos" gorm:"foreignKey:RestaurantID"`
    CreatedByID uint           `json:"created_by"`
    CreatedBy   User           `gorm:"foreignKey:CreatedByID" validate:"-"`
    Reviews     []Review       `json:"reviews"`
    Status      string         `gorm:"size:20" json:"status"` // "pending", "approved", "rejected"
}

// New input struct for creating a restaurant
type CreateRestaurantInput struct {
    Name        string `json:"name" validate:"required"`
    Address     string `json:"address" validate:"required"`
    Category    string `json:"category" validate:"required"`
    Description string `json:"description" validate:"required"`
}