package model

import "gorm.io/gorm"

type Station struct {
	gorm.Model
	LocationName string
}

type StationPiles struct {
	gorm.Model
	StationId uint
	PileId    uint
}
