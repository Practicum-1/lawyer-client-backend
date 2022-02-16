package seeds

import "github.com/Practicum-1/lawyer-client-backend.git/models"

func SeedPracticeArea() {
	DB.Model(&models.PracticeArea{}).Create(practice_area)
}

var practice_area []models.PracticeArea = []models.PracticeArea{
	{Name: "Personal"},
	{Name: "Divorce"},
	{Name: "Family Dispute"},
	{Name: "Child Custody"},
	{Name: "Muslim Law"},
	{Name: "Medical Negligence"},
	{Name: "Motor Accident"},
	{Name: "Wills"},
	{Name: "Trusts"},
	{Name: "Labour & Service"},
	{Name: "Corporate Law"},
	{Name: "Arbitration"},
	{Name: "Trademark & Copyright"},
	{Name: "Customs & Central Excise"},
	{Name: "Startup"},
	{Name: "Banking"},
	{Name: "Finance"},
	{Name: "GST"},
	{Name: "Corporate"},
	{Name: "Tax"},
	{Name: "Civil"},
	{Name: "Debt Matters"},
	{Name: "Documentation"},
	{Name: "Consumer Court"},
	{Name: "Civil"},
	{Name: "Cheque Bounce"},
	{Name: "Recovery"},
	{Name: "Property"},
	{Name: "Criminal"},
	{Name: "Landlord"},
	{Name: "Tenant"},
	{Name: "Cyber Crime"},
	{Name: "Others"},
	{Name: "Armed Forces Tribunal"},
	{Name: "Insurance"},
	{Name: "Immigration"},
	{Name: "International Law"},
}
