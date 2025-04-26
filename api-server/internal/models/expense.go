package models

import "time"

type Expense struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Amount      int `gorm:"not null" validate:"required"`
	Category     int `gorm:"not null" validate:"required"`
	PaidAt  time.Time `gorm:"not null;type:date;column:paid_at" validate:"required"`
	Description string `gorm:"not null" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID  int    `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" validate:"omitempty"`
}
