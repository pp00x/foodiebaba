package models

import (
    "gorm.io/gorm"
    "time"
)

type Review struct {
    ID           uint           `gorm:"primaryKey" json:"id"`
    CreatedAt    time.Time      `json:"created_at"`
    UpdatedAt    time.Time      `json:"updated_at"`
    DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    Rating       int            `json:"rating" validate:"required,min=1,max=5"`
    Comment      string         `gorm:"type:text" json:"comment" validate:"required"`
    UserID       uint           `json:"user_id"`
    User         User           `gorm:"foreignKey:UserID"`
    RestaurantID uint           `json:"restaurant_id" validate:"required"`
    Restaurant   Restaurant     `gorm:"foreignKey:RestaurantID"`
}