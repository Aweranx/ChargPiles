package model

import (
	"ChargPiles/consts"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId uint
	PileId uint
	Status string `gorm:"default:'unpaid'"`
}

func (order *Order) CompleteOrder() {
	order.Status = consts.Paid
}
