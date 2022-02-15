package models

import (
	"time"
)

type Lawyer struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"" json:"last_name"`
	// FullName      string         `gorm:"->;type:GENERATED ALWAYS AS (concat(first_name,' ',last_name));default:(-);"`
	Location         string                 `gorm:"not null" json:"location"`
	Email            string                 `gorm:"unique;not null" json:"email"`
	Password         string                 `gorm:"not null" json:"password"`
	ImageURL         string                 `gorm:"" json:"image_url"`
	Education        string                 `gorm:"not null" json:"education"`
	Experience       uint                   `gorm:"not null" json:"experience"`
	Gender           string                 `gorm:"not null" json:"gender"`
	Reviews          []Review               `gorm:"foreignKey:LawyerID" json:"reviews"`
	Courts           []Court                `gorm:"many2many:lawyer_courts" json:"courts"`
	Languages        []Language             `gorm:"many2many:lawyer_languages" json:"languages"`
	PracticeAreas    []LawyerPracticeArea   `gorm:"foreignKey:LawyerID" json:"practice_areas"`
	Phone            string                 `gorm:"not null" json:"phone"`
	PendingRequests  []Request              `gorm:"foreignKey:LawyerID;check: status === 'pending'" json:"pending_requests"` //just for testing
	Requests         []Request              `gorm:"foreignKey:LawyerID" json:"requests"`
	ConnectedClients []LawyerClientRelation `gorm:"foreignKey:LawyerID" json:"connected_lawyers"`
	CreatedAt        time.Time              `gorm:"" json:"created_at"`
	UpdatedAt        time.Time              `gorm:"" json:"updated_at"`
}
