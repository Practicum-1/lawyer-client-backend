package models

type Client struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"" json:"last_name"`
	// FullName  string   `gorm:"->;type:GENERATED ALWAYS AS (concat(firstname,' ',lastname));default:(-);"`
	Location         string                 `gorm:"not null" json:"location"`
	Email            string                 `gorm:"unique;not null" json:"email"`
	Phone            string                 `gorm:"not null" json:"phone"`
	Password         string                 `gorm:"not null" json:"password"`
	ImageURL         string                 `gorm:"" json:"image_url"`
	Reviews          []Review               `gorm:"foreignKey:ClientID" json:"reviews"`
	Requests         []Request              `gorm:"foreignKey:ClientID" json:"requests"`
	ConnectedLawyers []LawyerClientRelation `gorm:"foreignKey:ClientID" json:"connected_lawyers"`
}
