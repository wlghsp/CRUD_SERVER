package service

import (
	"CRUD_SERVER/repository"
	"CRUD_SERVER/types"
)

type User struct {
	userRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *User {
	return &User{
		userRepository: userRepository,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.userRepository.Create(newUser)
}

func (u *User) Update(beforeUser *types.User, updateUser *types.User) error {
	return u.userRepository.Update(beforeUser, updateUser)
}

func (u *User) Delete(newUser *types.User) error {
	return u.userRepository.Delete(newUser)
}

func (u *User) Get() []*types.User {
	return u.userRepository.Get()
}
