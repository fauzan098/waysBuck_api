package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	UserID    int       `json:"user_id" gorm:"type: int"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProductResponse struct {
	ID     int    `json:"id"`
	Title  string `json:"tittle"`
	Price  int    `json:"price"`
	Image  string `json:"image"`
	UserID int    `json:"user_id"`
}
