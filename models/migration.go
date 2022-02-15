package models

import (
	database "github.com/Practicum-1/lawyer-client-backend.git/db"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Migrate() {
	db := database.GetDB()
	DB = db
	// defer db.Migrator().DropTable(&Lawyer{}, &Client{}, &Review{}, &Court{}, &Language{}, &PracticeArea{}, "lawyer_courts", "lawyer_languages", "lawyer_practice-areas", &Request{})
	DB.AutoMigrate(&Court{}, &PracticeArea{}, &Language{}, &Client{}, &Lawyer{}, &LawyerClientRelation{}, &Chat{}, &LawyerPracticeArea{}, &Request{}, &Review{}) // Migrate the Book table.
}
