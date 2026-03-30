package services

import (
    "AuthInGo/db/repositories"
)


type UserService interface {
    CreateUser() error
}

type UserServiceImpl struct {
    userRepository repositories.UserRepository
}

func NewUserService(ur repositories.UserRepository) *UserServiceImpl {
    return &UserServiceImpl{
        userRepository: ur,
    }
}

func (u *UserServiceImpl) CreateUser() error {
    u.userRepository.Create()
    return nil
}