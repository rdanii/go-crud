package entity

type BookRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}

type BookRequestUpdate struct {
	ID          int    `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}
