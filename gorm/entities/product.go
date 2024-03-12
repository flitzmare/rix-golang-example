package entities

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint
	Name      string
	Price     float64
	UserID    uint
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   gorm.DeletedAt
}
