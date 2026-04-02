package services

import (
    db "AuthInGo/db/repositories"
    "fmt"
    "AuthInGo/utils"
)


type UserService interface {
    GetUserById() error
    CreateUser() error
    LoginUser() error
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

func (u *UserServiceImpl) CreateUser() error {

    fmt.Println("Creating user in UserService")

    // Example user details
    username := "john_doe"
    email := "john@example.com"
    password := "securepassword"

    // Hash the password before storing it in the database

    hashedPassword, err := utils.HashPassword(password)

    if err != nil {
        fmt.Println("Error hashing password:", err)
        return err
    }

    // Create the user in the database using the user repository

    u.userRepository.Create(
        username,
        email, 
        hashedPassword,
    )

    return nil
}

func (u *UserServiceImpl) LoginUser() error {
    response := utils.CheckPasswordHash("securepassword", "$2a$10$7a8b9c0d1e2f3g4h5i6j7k8l9m0n1o2p3q4r5s6t7u8v9w0x1y2z3")
    fmt.Println("Password match:", response)
    return nil
}

