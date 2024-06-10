package repository

import "CRUD_SERVER/types"

type UserRepository struct {
	UserMap []*types.User
}

func newUserRepository() *UserRepository {
	return &UserRepository{
		UserMap: []*types.User{},
	}
}

func (u *UserRepository) Create(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Update(beforeUser *types.User, updateUser *types.User) error {
	return nil
}

func (u *UserRepository) Delete(newUser *types.User) error {
	return nil
}

func (u *UserRepository) Get() []*types.User {
	return u.UserMap
}
