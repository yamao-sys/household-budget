package models

import (
	"time"
)

type Income struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Amount      int `gorm:"not null" validate:"required"`
	ReceivedAt  time.Time `gorm:"not null;type:date;column:received_at" validate:"required"`
	ClientName string `gorm:"not null" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID  int    `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" validate:"omitempty"`
}
