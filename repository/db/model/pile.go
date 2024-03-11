package model

import (
	"ChargPiles/consts"
	"gorm.io/gorm"
)

type Pile struct {
	gorm.Model
	PhysicisId uint   `gorm:"column:physicis_id;unique"`
	Status     string `gorm:"column:status;default:'available'"`
	StationId  uint   `gorm:"column:station_id;"`
}

// SetStatusInUse sets the status to "in use"
func (p *Pile) InUse() {
	p.Status = consts.InUse
}

// SetStatusReserved sets the status to "reserved"
func (p *Pile) Reserved() {
	p.Status = consts.Reserved
}

// SetStatusAvailable sets the status to "available"
func (p *Pile) Available() {
	p.Status = consts.Available
}

// SetStatusMaintenance sets the status to "maintenance"
func (p *Pile) Maintenance() {
	p.Status = consts.Maintenance
}

func (p *Pile) ChangeBase(physicisId uint) {
	p.PhysicisId = physicisId
}

func (p *Pile) ChangeStation(stationId uint) {
	p.StationId = stationId
}

func CreatePile(physicisId uint, stationId uint) (pile *Pile) {
	pile.PhysicisId = physicisId
	pile.StationId = stationId
	return pile
}
