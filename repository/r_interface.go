package repository

import (
	"errors"
	"go-crud/entity"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(reqBook entity.BookRequest) (entity.BookResponse, error)
	GetBooks() ([]entity.BookResponse, error)
	GetBookByID(id int) (entity.BookResponse, error)
	Update(reqBook entity.BookRequestUpdate) (entity.BookResponse, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) Create(reqBook entity.BookRequest) (res entity.BookResponse, err error) {

	var existingBook entity.Book

	err = r.db.Debug().Table(entity.TABLE_BOOKS).Where("title = ? AND description = ? AND price = ? AND rating = ?", reqBook.Title, reqBook.Description, reqBook.Price, reqBook.Rating).First(&existingBook).Error
	if err == nil {
		return res, errors.New("data already exists")
	}

	data := entity.Book{
		Title:       reqBook.Title,
		Description: reqBook.Description,
		Price:       reqBook.Price,
		Rating:      reqBook.Rating,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = r.db.Debug().Table(entity.TABLE_BOOKS).Create(&data).Error
	if err != nil {
		return res, err
	}

	bookResponse := entity.BookResponse{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Price:       data.Price,
		Rating:      data.Rating,
	}

	return bookResponse, nil
}

func (r *repository) GetBooks() (res []entity.BookResponse, err error) {
	err = r.db.Debug().Table(entity.TABLE_BOOKS).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repository) GetBookByID(id int) (res entity.BookResponse, err error) {
	err = r.db.Debug().Table(entity.TABLE_BOOKS).First(&res, id).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r *repository) Update(reqBook entity.BookRequestUpdate) (entity.BookResponse, error) {
	var data entity.Book
	var res entity.BookResponse

	// Menggunakan Find untuk mengambil data buku berdasarkan ID
	result := r.db.Debug().Table(entity.TABLE_BOOKS).Find(&data, "id = ?", reqBook.ID)
	if result.RowsAffected == 0 {
		return res, errors.New("not found")
	}

	// Memeriksa apakah data sudah sama dengan data yang ada di database
	if reqBook.Title == data.Title && reqBook.Price == data.Price && reqBook.Description == data.Description && reqBook.Rating == data.Rating {
		return res, errors.New("data already exists")
	}

	// Memperbarui data buku
	data.Title = reqBook.Title
	data.Description = reqBook.Description
	data.Price = reqBook.Price
	data.Rating = reqBook.Rating

	// Memperbarui waktu jika data baru
	if data.CreatedAt.IsZero() {
		data.CreatedAt = time.Now()
	}
	data.UpdatedAt = time.Now()

	// Menggunakan Save untuk menyimpan perubahan data ke dalam database
	err := r.db.Debug().Table(entity.TABLE_BOOKS).Save(&data).Error
	if err != nil {
		return res, err
	}

	// Menyiapkan response data
	dataResponse := entity.BookResponse{
		ID:          reqBook.ID,
		Title:       data.Title,
		Description: data.Description,
		Price:       data.Price,
		Rating:      data.Rating,
	}

	return dataResponse, nil
}

func (r *repository) Delete(id int) error {
	var data entity.Book

	err := r.db.Debug().Table(entity.TABLE_BOOKS).Where("id = ?", id).First(&data).Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}
