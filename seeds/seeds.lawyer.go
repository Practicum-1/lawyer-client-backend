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
		Location:   "Baltimore",
		Email:      "alie@gmail.com",
		Phone:      "+912211",
		Password:   "lawyer123",
		Gender:     "female",
		Education:  "LLM",
		Experience: 5,
		// Languages: []models.Language{
		// 	{Name: "Assamese"},
		// 	{Name: "Bengali"},
		// },
	},
	{
		FirstName:  "Lian",
		LastName:   "Vade",
		Location:   "Asgard",
		Email:      "lian@gmail.com",
		Phone:      "+911123123",
		Password:   "lawyer123",
		Gender:     "male",
		Education:  "LLB",
		Experience: 10,
	},
	{
		FirstName:  "John",
		LastName:   "Jones",
		Location:   "London",
		Email:      "john@gmail.com",
		Phone:      "+912321321",
		Password:   "lawyer123",
		Gender:     "female",
		Education:  "Integrated LLB",
		Experience: 12,
	},
}
