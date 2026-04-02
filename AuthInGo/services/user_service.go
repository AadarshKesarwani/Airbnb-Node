package services

import (
    db "AuthInGo/db/repositories"
)


type UserService interface {
    CreateUser() error
}

type UserServiceImpl struct {
    userRepository db.UserRepository
}

func NewUserService(ur db.UserRepository) *UserServiceImpl {
    return &UserServiceImpl{
        userRepository: ur,
    }
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Fetching user in UserService")
	u.userRepository.GetByID()
	return nil
}