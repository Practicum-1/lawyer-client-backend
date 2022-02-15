package repositories

import (
	"errors"

	"github.com/Practicum-1/lawyer-client-backend.git/db"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"gorm.io/gorm"
)

func GetClientById(id uint64, client *models.Client) error {
	db := db.GetDB()
	result := db.Where("id = ?", id).First(&client)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("404")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
