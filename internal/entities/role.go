package entities

import "github.com/mileapp_screening/internal/entities/base"

type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	base.Timestamp
}
