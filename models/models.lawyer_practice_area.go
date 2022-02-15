package models

import "time"

type LawyerPracticeArea struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	LawyerID       uint      `gorm:"not null" json:"lawyer_id"`
	PracticeAreaID uint      `gorm:"not null" json:"practice_area_id"`
	Charge         uint      `gorm:"not null" json:"charge"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
