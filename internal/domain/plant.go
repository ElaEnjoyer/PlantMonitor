package domain

import "time"

type Plant struct {
	Id        uint64
	UserId    uint64
	Name      string
	City      string
	Address   string
	Type      PlantType
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type PlantType string

const (
	SolarType PlantType = "SOLAR"
	WindType  PlantType = "WIND"
)
