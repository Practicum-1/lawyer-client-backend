package repositories

import (
	"errors"

	"github.com/Practicum-1/lawyer-client-backend.git/db"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"gorm.io/gorm"
)

func GetAllClients() (interface{}, error) {
	db := db.GetDB()
	var client []models.Client
	result := db.Find(&client)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("404")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return client, nil
}

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

func CreateClient(user *models.Client) error {
	db := db.GetDB()
	if result := db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}
