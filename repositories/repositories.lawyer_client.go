package repositories

import (
	"errors"

	"github.com/Practicum-1/lawyer-client-backend.git/db"
	"github.com/Practicum-1/lawyer-client-backend.git/models"
	"gorm.io/gorm"
)

func CreateLawyerClientConnection(lawyerID, clientID uint) error {
	db := db.GetDB()
	lawyerClient := &models.LawyerClient{
		LawyerID: lawyerID,
		ClientID: clientID,
	}
	result := db.Where("lawyer_id = ?", lawyerID).Where("client_id = ?", clientID).FirstOrCreate(&lawyerClient)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		lawyerClient = &models.LawyerClient{
			LawyerID: lawyerID,
			ClientID: clientID,
		}
		result = db.Create(lawyerClient)
		if result.Error != nil {
			return result.Error
		}
		return nil
	} else if result.Error != nil {
		return result.Error
	} else {
		return nil
	}
}

func GetAllClientsByLawyerID(lawyerID string) ([]models.LawyerClient, error) {
	db := db.GetDB()
	var lawyerClients *[]models.LawyerClient
	// var clients []*models.Client
	result := db.Model(&models.LawyerClient{}).Preload("Client").Where("lawyer_id = ?", lawyerID).Find(&lawyerClients)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("404")
	} else if result.Error != nil {
		return nil, result.Error
	} else {
		// for _, lawyerClient := range *lawyerClients {
		// 	clients = append(clients, lawyerClient.Client)
		// }
		return *lawyerClients, nil
	}
}

func GetAllLawyersByClientID(clientID string) ([]models.LawyerClient, error) {
	db := db.GetDB()
	var lawyerClients *[]models.LawyerClient
	// var lawyers []*models.Lawyer
	result := db.Model(&models.LawyerClient{}).Preload("Client").Where("lawyer_id = ?", clientID).Find(&lawyerClients)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("404")
	} else if result.Error != nil {
		return nil, result.Error
	} else {
		// for _, lawyerClient := range *lawyerClients {
		// 	lawyers = append(lawyers, lawyerClient.Lawyer)
		// }
		return *lawyerClients, nil
	}
}

func CreateChat(chat *models.Chat) error {
	db := db.GetDB()
	result := db.Model(&models.Chat{}).Create(chat)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLawyerClientChat(lc string) ([]models.Chat, error) {
	db := db.GetDB()
	var chat []models.Chat
	result := db.Where("lawyer_client_id = ?", lc).Order("created_at desc").Find(&chat)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("404")
	} else if result.Error != nil {
		return nil, result.Error
	} else {
		return chat, nil
	}
}
