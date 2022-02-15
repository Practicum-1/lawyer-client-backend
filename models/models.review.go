package models

import "time"

type Review struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	LawyerID    uint      `gorm:"not null" json:"lawyer_id"`
	Lawyer      Lawyer    `gorm:"foreignKey:LawyerID" json:"lawyer"`
	ClientID    uint      `gorm:"not null" json:"client_id"`
	Client      Client    `gorm:"foreignKey:ClientID" json:"client"`
	Description string    `gorm:"" json:"description"`
	Rating      uint      `gorm:"" json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
