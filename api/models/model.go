package models

import (
	"time"
)

type Todo struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	createat time.Time
	updateat time.Time
}
