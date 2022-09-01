package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int64          `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Quantity    int            `json:"quantity"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type CreateProductInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int `json:"quantity"`
}
