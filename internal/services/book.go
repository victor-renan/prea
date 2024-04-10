package services

import (
	"prea/internal/domain/models"
	genericRepos "prea/internal/repositories/generic"
)

type IBookService interface {
	GetAll() ([]models.Book, error)
	GetById(id int64) (models.Book, error)
	Create(entity models.Book) (models.Book, error)
	Update(id int64, partial models.Book) (models.Book, error)
	Delete(id int64) error
}

type BookService struct {
	Repo genericRepos.IGenericRepository[models.Book]
}

func (bs BookService) GetAll() ([]models.Book, error) {
	return bs.Repo.GetAll()
}

func (bs BookService) GetById(id int64) (models.Book, error) {
	return bs.Repo.GetById(id)
}

func (bs BookService) Create(entity models.Book) (models.Book, error) {
	return bs.Repo.Create(entity)
}

func (bs BookService) Update(id int64, partial models.Book) (models.Book, error) {
	return bs.Repo.Update(id, partial)
}

func (bs BookService) Delete(id int64) error {
	return bs.Repo.Delete(id)
}
