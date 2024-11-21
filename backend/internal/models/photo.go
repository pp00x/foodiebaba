package models

import (
    "gorm.io/gorm"
    "time"
)

type Photo struct {
    ID          uint           `gorm:"primaryKey" json:"id"`
    CreatedAt   time.Time      `json:"created_at"`
    UpdatedAt   time.Time      `json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
    URL         string         `gorm:"size:255" json:"url"`
    RestaurantID uint          `json:"restaurant_id"`
}