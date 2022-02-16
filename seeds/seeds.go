package seeds

import (
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func RunSeeds() {
	DB = models.DB
	SeedCourts()
	SeedLanguages()
	SeedPracticeArea()
}
