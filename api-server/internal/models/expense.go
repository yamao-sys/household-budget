package models

import (
	"time"
)

type Category int

const (
    CategoryFood Category = iota + 1
    CategoryDailyGoods
    CategoryExtra
	CategoryEntertainment
	CategorySelfInvestment
	CategoryLiving
	CategoryMiscellaneous
	CategoryTaxSaving
	CategorySavings
)

type Expense struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Amount      int `gorm:"not null" validate:"required"`
	Category     Category `gorm:"not null" validate:"required"`
	PaidAt  time.Time `gorm:"not null;type:date;column:paid_at" validate:"required"`
	Description string `gorm:"not null" validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID  int    `gorm:"not null" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID" validate:"omitempty"`
}

// func (s Category) Value() (driver.Value, error) {
//     return int64(s), nil
// }

// func (s *Category) Scan(value interface{}) error {
//     intVal, ok := value.(int64)
//     if !ok {
//         return fmt.Errorf("failed to scan Status: %v", value)
//     }
//     *s = Category(intVal)
//     return nil
// }

// func (c Category) String() string {
//     switch c {
//     case CategoryFood:
//         return "food"
//     case CategoryDailyGoods:
//         return "daily_goods"
//     case CategoryExtra:
//         return "extra""
//     case CategoryEntertainment
//         return "entertainment"
//     case CategorySelfInvestment:
//         return "self_investment"
//     case CategoryLiving:
//         return "living"
//     case CategoryMiscellaneous:
//         return "miscellaneous"
//     case CategoryTaxSaving:
//         return "tax_saving"
//     case CategorySavings:
//         return "savings"
//     default:
//         return "unknown"
//     }
// }
