package models

import (
	"errors"

	"github.com/Practicum-1/lawyer-client-backend.git/helpers"
	"gorm.io/gorm"
)

type Client struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"" json:"last_name"`
	// FullName  string   `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
	Location         string                 `gorm:"not null" json:"location"`
	Email            string                 `gorm:"unique;not null" json:"email"`
	Phone            string                 `gorm:"not null" json:"phone"`
	Password         string                 `gorm:"not null" json:"password"`
	ImageURL         string                 `gorm:"" json:"image_url"`
	Reviews          []Review               `gorm:"foreignKey:ClientID" json:"reviews"`
	Requests         []Request              `gorm:"foreignKey:ClientID" json:"requests"`
	ConnectedLawyers []LawyerClientRelation `gorm:"foreignKey:ClientID" json:"connected_lawyers"`
}

func (client *Client) BeforeCreate(tx *gorm.DB) error { //Validate the client before creating it
	var count int64
	DB.Model(&client).Where("email = ? OR phone = ?", client.Email, client.Phone).Count(&count)
	if count != 0 {
		return errors.New("user already exists with this email or phone number")
	}
	hashedPassword, err := helpers.HashPassword(client.Password)
	if err != nil {
		return err
	}
	client.Password = hashedPassword
	return nil
}
