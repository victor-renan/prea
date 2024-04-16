package user

import (
	"prea/internal/generics/repositories"
	"prea/internal/security"
	"time"
)

type IUserService interface {
	GetAll() ([]User, error)
	GetById(id string) (User, error)
	GetByUsername(username string) (User, error)
	Create(entity *UserCreateDAO) (User, error)
	Update(id string, partial *UserUpdateDAO) (User, error)
	AlterLastLogin(user User, lastlogin time.Time) (*time.Time, *time.Time, error)
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

func (bs UserService) GetByUsername(username string) (User, error) {
	return bs.Repo.GetFirst(User{Username: username})
}

func (bs UserService) Create(entity *UserCreateDAO) (User, error) {
	pw, err := security.HashPassword(entity.Password)
	if err != nil {
		return User{}, err
	}

	entity.Password = pw

	return bs.Repo.Create(&entity)
}

func (bs UserService) Update(id string, partial *UserUpdateDAO) (User, error) {
	return bs.Repo.Update(id, &partial)
}

func (bs UserService) AlterLastLogin(
	user User,
	lastlogin time.Time,
) (new *time.Time, old *time.Time, err error) {
	old = user.LastLogin
	updated, err := bs.Repo.Update(user.Id, struct{ LastLogin *time.Time }{&lastlogin})
	new = updated.LastLogin

	return
}

func (bs UserService) Delete(id string) error {
	return bs.Repo.Delete(id)
}
