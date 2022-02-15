package models

import "time"

type RequestData struct {
	Title           string `json:"title"`
	Body            string `json:"body"`
	DenyWithMessage string `json:"deny_with_message"`
}

type Request struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	LawyerID    uint        `gorm:"not null" json:"lawyer_id"`
	ClientID    uint        `gorm:"not null" json:"client_id"`
	RequestData RequestData `gorm:"embedded" json:"request_data"`
	Status      string      `gorm:"not null" json:"status"`
	CreatedAt   time.Time   `gorm:"" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"" json:"updated_at"`
}
