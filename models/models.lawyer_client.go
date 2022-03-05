package models

import "time"

type LawyerClient struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	LawyerID    uint      `gorm:"not null" json:"lawyer_id"`
	Lawyer      Lawyer    `gorm:"foreignKey:LawyerID" json:"lawyer"`
	ClientID    uint      `gorm:"not null" json:"client_id"`
	Client      Client    `gorm:"foreignKey:client_id" json:"client"`
	Chats       []Chat    `gorm:"foreignKey:LawyerClientID" json:"chats"`
	TagByLawyer string    `gorm:"not null" json:"lawyer_tag"`
	TagByClient string    `gorm:"not null" json:"client_tag"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
