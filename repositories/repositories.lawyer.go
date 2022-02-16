package repositories

import (
	"github.com/Practicum-1/lawyer-client-backend.git/db"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
)

func CreateLayer(lawyer *models.Lawyer) error {
	db := db.GetDB()
	if err := db.Create(&lawyer).Error; err != nil {
		return err
	}
	return nil
}
