package usecase

import (
	"homework.31/pkg/entity"
)

type (
	Usecase interface {
		CreateUser(*entity.User) (int64, error)
		MakeFriends(*entity.MakeFriends) (int, int, error)
		DeleteUser(user *entity.DeleteUser) string
		UpdateAge(user *entity.UpdateUser) string
		GetFriends(int) (string, error)
	}

	Repository interface {
		CreateUser(*entity.User) (int64, error)
		DeleteUser(user *entity.DeleteUser) string
		UpdateAge(user *entity.UpdateUser) string
		MakeFriends(*entity.MakeFriends) (int, int, error)
		GetFriends(int) (string, error)
	}
)

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) *usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) CreateUser(user *entity.User) (int64, error) {
	uid, error := u.repository.CreateUser(user)
	return uid, error
}

func (u *usecase) MakeFriends(friends *entity.MakeFriends) (a, b int, err error) {
	a, b, err = u.repository.MakeFriends(friends)
	return a, b, err
}

func (u *usecase) DeleteUser(user *entity.DeleteUser) string {
	b := u.repository.DeleteUser(user)
	return b
}

func (u *usecase) GetFriends(a int) (b string, err error) {
	b, err = u.repository.GetFriends(a)
	return b, err
}

func (u *usecase) UpdateAge(user *entity.UpdateUser) string {
	s := u.repository.UpdateAge(user)
	return s
}
