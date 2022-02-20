package seeds

import "github.com/Practicum-1/lawyer-client-backend.git/models"

func SeedClients() {
	var count int64
	DB.Model(&models.Client{}).Count(&count)
	if count == 0 {
		DB.Model(&models.Client{}).Create(clients)
	}
}

var clients = []models.Client{
	{
		FirstName: "Jason",
		LastName:  "Voorhees",
		Location:  "Baltimore",
		Email:     "jason@gmail.com",
		Phone:     "+912",
		Password:  "client123"},
	{
		FirstName: "Scorpian",
		LastName:  "Vade",
		Location:  "Asgard",
		Email:     "scorpian@gmail.com",
		Phone:     "+911",
		Password:  "client123"},
	{
		FirstName: "Miranda",
		LastName:  "Jones",
		Location:  "London",
		Email:     "miranda@gmail.com",
		Phone:     "+912",
		Password:  "client123"},
}
