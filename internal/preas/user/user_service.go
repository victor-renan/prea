package user

import "prea/internal/generics/repositories"

type IUserService interface {
	GetAll() ([]User, error)
	GetById(id string) (User, error)
	Create(entity User) (User, error)
	Update(id string, partial User) (User, error)
	Delete(id string) error
}

type UserService struct {
	Repo repositories.IGenericRepository[User]
}

func (bs UserService) GetAll() ([]User, error) {
	return bs.Repo.GetAll()
}

func (bs UserService) GetById(id string) (User, error) {
	return bs.Repo.GetById(id)
}

func (bs UserService) Create(entity User) (User, error) {
	return bs.Repo.Create(entity)
}

func (bs UserService) Update(id string, partial User) (User, error) {
	return bs.Repo.Update(id, partial)
}

func (bs UserService) Delete(id string) error {
	return bs.Repo.Delete(id)
}
