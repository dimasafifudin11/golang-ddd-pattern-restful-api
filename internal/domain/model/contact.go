package model

import "time"

// Contact represents the contact entity.
type Contact struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	FirstName string    `gorm:"not null" json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Addresses []Address `gorm:"foreignKey:ContactID" json:"addresses,omitempty"` // Relasi one-to-many
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
