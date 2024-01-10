package entity

const (
	TABLE_BOOKS = "books"
)

type Tabler interface {
	TableName() string
}

func (Book) TableName() string {
	return TABLE_BOOKS
}
