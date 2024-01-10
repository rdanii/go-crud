package usecase

import (
	"go-crud/entity"
	"go-crud/repository"
)

type Usecase interface {
	Create(reqBook entity.BookRequest) (entity.BookResponse, error)
	GetBooks() ([]entity.BookResponse, error)
	GetBookByID(id int) (entity.BookResponse, error)
	Update(reqBook entity.BookRequestUpdate) (entity.BookResponse, error)
	Delete(id int) error
}

type usecase struct {
	repository repository.Repository
}

func NewUsecase(repository repository.Repository) Usecase {
	return &usecase{repository}
}

func (u *usecase) Create(reqBook entity.BookRequest) (entity.BookResponse, error) {
	book, err := u.repository.Create(reqBook)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (u *usecase) GetBooks() ([]entity.BookResponse, error) {
	book, err := u.repository.GetBooks()
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (u *usecase) GetBookByID(id int) (entity.BookResponse, error) {
	book, err := u.repository.GetBookByID(id)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (u *usecase) Update(reqBook entity.BookRequestUpdate) (entity.BookResponse, error) {
	book, err := u.repository.Update(reqBook)
	if err != nil {
		return book, err
	}

	return book, nil
}

func (u *usecase) Delete(id int) error {
	err := u.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
