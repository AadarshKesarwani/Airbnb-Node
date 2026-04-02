package services

import (
    env "AuthInGo/config/env"
    db "AuthInGo/db/repositories"
    "AuthInGo/models"
    "AuthInGo/dto"
    "AuthInGo/utils"
    "fmt"
    "github.com/golang-jwt/jwt/v5"
)




type UserService interface {
    GetUserById(id int64) (*models.User, error)
    CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error)
    LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
}




type UserServiceImpl struct {
    userRepository db.UserRepository
}





func NewUserService(ur db.UserRepository) *UserServiceImpl {
    return &UserServiceImpl{
        userRepository: ur,
    }
}




func (u *UserServiceImpl) GetUserById(id int64) (*models.User, error) {
	fmt.Println("Fetching user in UserService")
	user, err := u.userRepository.GetByID(id)

    if err != nil {
        fmt.Println("Error fetching user by ID:", err)
        return nil, err
    }

    return user, nil
}



func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDTO) (*models.User, error) {

    fmt.Println("Creating user in UserService")

    // Example user details
    username := payload.Username
    email := payload.Email
    password := payload.Password

    // Hash the password before storing it in the database

    hashedPassword, err := utils.HashPassword(password)

    if err != nil {
        fmt.Println("Error hashing password:", err)
        return nil, err
    }

    // Create the user in the database using the user repository

    user, err := u.userRepository.Create(
        username,
        email, 
        hashedPassword,
    )

    if err != nil {
        fmt.Println("Error creating user:", err)
        return nil, err
    }

    return user, nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
    fmt.Println("Logging in user in UserService")

    // Example user details for login
    email := payload.Email
    password := payload.Password

    //step 1: Fetch the user from the database using the email

    user, err := u.userRepository.GetByEmail(email)

    //step 2: If user is not found, return an error
    if err != nil {
        fmt.Println("Error fetching user by email:", err)
        return "", err
    }
    if user == nil {
        fmt.Println("No user found with email:", email)
        return "", fmt.Errorf("invalid email or password")
    }

    //step 3: Compare the provided password with the hashed password stored in the database

    isPasswordValid := utils.CheckPasswordHash(password, user.Password)

    if !isPasswordValid {
        fmt.Println("Invalid password for email:", email)
        return "", fmt.Errorf("invalid email or password")
    }

    //step 4: If the password is valid, generate a JWT token for the user

    jwtPayload := jwt.MapClaims{
        "id": user.Id,
        "email":   user.Email,
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

    // Sign the token with a secret key (you should use a secure key in production)

    secretKey := []byte(env.GetString("JWT_SECRET_KEY", "your_secret_key_here"))

    // Generate the signed token string

    tokenString, err := token.SignedString(secretKey)

    if err != nil {
        fmt.Println("Error generating JWT token:", err)
        return "", err
    }

    fmt.Println("User logged in successfully, generated JWT token:", tokenString)

    return tokenString, nil

}

