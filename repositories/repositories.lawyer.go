package repositories

import (
	"errors"
	"fmt"
	"strconv"

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

func getIDs(lawyers []models.Lawyer) []uint {
	var resultIDs []uint
	for _, lawyer := range lawyers {
		resultIDs = append(resultIDs, lawyer.ID)
	}
	return resultIDs
}

func GetLawyerByFilter(filters models.Filters) ([]models.Lawyer, error) {
	db := db.GetDB()
	var lawyers []models.Lawyer
	var noOfAssociationFilters int = 0
	var statement string = "SELECT * FROM lawyers WHERE 1=1"
	if filters.Experience != 0 {
		statement += " AND `experience` = " + strconv.Itoa(int(filters.Experience))
	}
	if filters.LocationID != 0 {
		statement += " AND `location_id` = " + strconv.Itoa(int(filters.LocationID))
	}
	if filters.Gender != "" {
		statement += " AND `gender` = \"" + filters.Gender + "\""
	}
	if filters.PracticeAreaID != 0 {
		colName := func() string {
			if noOfAssociationFilters == 0 {
				return " `id`"
			} else {
				return " `lawyer_id`"
			}
		}()
		statement += " AND" + colName + " IN ( SELECT `lawyer_id` FROM lawyer_practice_areas WHERE practice_area_id = " + strconv.Itoa(int(filters.PracticeAreaID))
		noOfAssociationFilters += 1
	}
	if filters.CourtID != 0 {
		colName := func() string {
			if noOfAssociationFilters == 0 {
				return "`id`"
			} else {
				return "`lawyer_id`"
			}
		}()
		statement += " AND" + colName + " IN ( SELECT `lawyer_id` FROM lawyer_courts WHERE court_id = " + strconv.Itoa(int(filters.CourtID))
		noOfAssociationFilters += 1
	}
	if filters.LanguageID != 0 {
		colName := func() string {
			if noOfAssociationFilters == 0 {
				return "`id`"
			} else {
				return "`lawyer_id`"
			}
		}()
		statement += " AND" + colName + " IN ( SELECT `lawyer_id` FROM lawyer_languages WHERE language_id = " + strconv.Itoa(int(filters.LanguageID))
		noOfAssociationFilters += 1
	}
	for i := 0; i < noOfAssociationFilters; i++ {
		statement += ")"
	}
	statement += ";"
	fmt.Println(statement)
	result := db.Raw(statement).Find(&lawyers)
	fmt.Println("RowsAffected", result.RowsAffected)

	return lawyers, nil
}
