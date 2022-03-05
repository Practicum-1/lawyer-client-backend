package models

import (
	"gorm.io/gorm"
)

type Court struct {
	ID       uint   `gorm:"primaryKey" json:"court_id"`
	Name     string `gorm:"not null" json:"name"`
	Type     string `gorm:"not null" json:"type"`
	Location string `gorm:"not null" json:"location"`
	gorm.Model
}

type Language struct {
	ID   uint   `gorm:"primaryKey" json:"language_id"`
	Name string `gorm:"not null" json:"name"`
	gorm.Model
}

type PracticeArea struct {
	ID     uint   `gorm:"primaryKey" json:"practice_area_id"`
	Name   string `gorm:"not null" json:"name"`
	AvgFee string `gorm:"not null" json:"avg_fee"`
	gorm.Model
}

type Location struct {
	ID    uint   `gorm:"primaryKey" json:"location_id"`
	City  string `gorm:"not null" json:"city"`
	State string `gorm:"not null" json:"state"`
}
