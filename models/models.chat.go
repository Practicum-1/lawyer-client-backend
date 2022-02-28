package models

import "time"

type Chat struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	LawyerClientID uint      `gorm:"not null" json:"lawyer_client_id"`
	Message        string    `gorm:"not null" json:"message"`
	SentByLawyer   bool      `gorm:"not null" json:"sent_by_lawyer"`
	ReplyToID      uint      `gorm:"default:null" json:"reply_to_id"`
	RepliedTo      *Chat     `gorm:"foreignKey:ReplyToID" json:"replied_to"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
