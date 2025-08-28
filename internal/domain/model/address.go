package model

// Address represents the address entity.
type Address struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ContactID  uint   `gorm:"not null" json:"contact_id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	Country    string `gorm:"not null" json:"country"`
	PostalCode string `json:"postal_code"`
}
