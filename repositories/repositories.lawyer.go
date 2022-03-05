package repositories

import (
	"errors"

	"github.com/Practicum-1/lawyer-client-backend.git/db"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"gorm.io/gorm"
)

func GetAllLawyers() (interface{}, error) {
	db := db.GetDB()
	var lawyer []models.Lawyer
	result := db.Find(&lawyer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("404")
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return lawyer, nil
}

func GetLawyerById(id uint64, lawyer *models.Lawyer) error {
	db := db.GetDB()
	// result := db.Preload(clause.Associations).Where("id = ?", id).First(&lawyer)
	result := db.Where("id = ?", id).First(&lawyer)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("404")
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateLawyer(lawyer *models.Lawyer) error {
	db := db.GetDB()
	if result := db.Create(&lawyer); result.Error != nil {
		return result.Error
	}
	return nil
}
