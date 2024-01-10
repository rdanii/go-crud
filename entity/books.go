package entity

import "time"

type Book struct {
	ID          int       `json:"id" gorm:"column:id"`
	Title       string    `json:"title" gorm:"column:title"`
	Description string    `json:"description" gorm:"column:description"`
	Price       int       `json:"price" gorm:"column:price"`
	Rating      int       `json:"rating" gorm:"column:rating"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at"`
}
