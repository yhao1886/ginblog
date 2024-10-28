package model

import "time"

type Model struct {
	ID        int       `gorm:"primary_key;auto_increament" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
