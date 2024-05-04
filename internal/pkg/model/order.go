package model

import "time"

type Orders []*Order

type Order struct {
	ID        string     `gorm:"column:id"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	UserID    string     `gorm:"column:user_id"`
	Payload   string     `gorm:"column:payload"`
	IsSuccess bool       `gorm:"column:is_success"`
}

func (*Order) TableName() string {
	return "orders"
}
