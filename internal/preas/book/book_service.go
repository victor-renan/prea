package book

import "prea/internal/generics/repositories"

type IBookService interface {
	GetAll() ([]Book, error)
	GetById(id int64) (Book, error)
	Create(entity Book) (Book, error)
	Update(id int64, partial Book) (Book, error)
	Delete(id int64) error
}

type BookService struct {
	Repo repositories.IGenericRepository[Book]
}

func (bs BookService) GetAll() ([]Book, error) {
	return bs.Repo.GetAll()
}

func (bs BookService) GetById(id int64) (Book, error) {
	return bs.Repo.GetById(id)
}

func (bs BookService) Create(entity Book) (Book, error) {
	return bs.Repo.Create(entity)
}

func (bs BookService) Update(id int64, partial Book) (Book, error) {
	return bs.Repo.Update(id, partial)
}

func (bs BookService) Delete(id int64) error {
	return bs.Repo.Delete(id)
}
