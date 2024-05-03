package model

import "time"

type Corrections []*Correction

type Correction struct {
	ID        string     `gorm:"column:id"`
	CretedAt  time.Time  `gorm:"column:creted_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	Payload   string     `gorm:"column:payload"`
	IsError   bool       `gorm:"column:is_error"`
}

func (*Correction) TableName() string {
	return "corrections"
}
