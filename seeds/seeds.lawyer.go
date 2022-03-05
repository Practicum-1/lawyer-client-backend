package seeds

import "github.com/Practicum-1/lawyer-client-backend.git/models"

func SeedLawyer() {
	var count int64
	DB.Model(&models.Lawyer{}).Count(&count)
	if count == 0 {
		DB.Model(&models.Lawyer{}).Create(lawyer)
	}
}

var lawyer = []models.Lawyer{
	{
		FirstName:  "Alie",
		LastName:   "Voorhees",
		LocationID: 4,
		Email:      "alie@gmail.com",
		Phone:      "+912211",
		Password:   "lawyer123",
		Gender:     "female",
		Education:  "LLM",
		Experience: 5,
		Languages: []models.Language{
			{ID: 1},
			{ID: 3},
			{ID: 5},
		},
		Courts: []models.Court{
			{ID: 3},
			{ID: 5},
			{ID: 10},
			{ID: 23},
		},
		PracticeAreas: []models.LawyerPracticeArea{
			{ID: 1},
			{ID: 2},
			{ID: 3},
			{ID: 4},
			{ID: 10},
			{ID: 13},
			{ID: 15},
			{ID: 17},
			{ID: 19},
			{ID: 22},
			{ID: 24},
			{ID: 25},
			{ID: 29},
			{ID: 31},
			{ID: 37},
		},
	},
	{
		FirstName:  "Lian",
		LastName:   "Vade",
		LocationID: 7,
		Email:      "lian@gmail.com",
		Phone:      "+911123123",
		Password:   "lawyer123",
		Gender:     "male",
		Education:  "LLB",
		Experience: 10,
		Languages: []models.Language{
			{ID: 15},
			{ID: 21},
			{ID: 19},
		},
		Courts: []models.Court{
			{ID: 6},
			{ID: 8},
			{ID: 13},
			{ID: 18},
		},
		PracticeAreas: []models.LawyerPracticeArea{
			{ID: 2},
			{ID: 3},
			{ID: 4},
			{ID: 5},
			{ID: 11},
			{ID: 14},
			{ID: 16},
			{ID: 18},
			{ID: 21},
			{ID: 24},
			{ID: 26},
			{ID: 30},
			{ID: 31},
			{ID: 32},
			{ID: 33},
		},
	},
	{
		FirstName:  "John",
		LastName:   "Jones",
		LocationID: 8,
		Email:      "john@gmail.com",
		Phone:      "+912321321",
		Password:   "lawyer123",
		Gender:     "female",
		Education:  "Integrated LLB",
		Experience: 12,
		Languages: []models.Language{
			{ID: 17},
			{ID: 13},
			{ID: 14},
		},
		Courts: []models.Court{
			{ID: 4},
			{ID: 7},
			{ID: 13},
			{ID: 25},
		},
		PracticeAreas: []models.LawyerPracticeArea{
			{ID: 3},
			{ID: 5},
			{ID: 8},
			{ID: 12},
			{ID: 13},
			{ID: 15},
			{ID: 16},
			{ID: 17},
			{ID: 18},
			{ID: 21},
			{ID: 23},
			{ID: 32},
			{ID: 33},
			{ID: 35},
			{ID: 36},
		},
	},
}
