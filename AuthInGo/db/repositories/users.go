package db

import (
    "AuthInGo/models"
    "database/sql"
    "fmt"
)

type UserRepository interface {
    Create(username string, email string, password string) (*models.User, error)
    GetByID(id int64) (*models.User, error)
    GetALL() ([]*models.User, error)
    DeleteById(id int64) error
    GetByEmail(email string) (*models.User, error)
}

type UserRepositoryImpl struct {
    db *sql.DB
}

func NewUserRepository(_db *sql.DB) *UserRepositoryImpl {
    return &UserRepositoryImpl{
        db: _db,
    }
}



func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) (*models.User, error) {

    // Implement the logic to create a new user in the database
    //means we are defining the SQL query to insert a new user record into the users table. The query uses placeholders (?) for the values that will be provided when executing the query. This allows us to safely insert user data into the database without risking SQL injection attacks.
    query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

    // Execute the query with the user details
    //means we are executing the SQL query using the Exec method of the database connection. We pass the user details (username, email, password) as arguments to the Exec method, which will replace the placeholders in the query with the actual values. The result variable will hold the result of the query execution, which can be used to check if the user was created successfully
    


    result, err := u.db.Exec(query, username, email, hashedPassword)



    // Check for errors during query execution
    //means we are checking if there was an error during the execution of the query. If there is an error, we log it and return the error to the caller. This allows us to handle any issues that may arise during the user creation process, such as database connection problems or constraints violations.


    if err != nil {
        fmt.Println("Error creating user:", err)
        return nil, err
    }

    // Check if the user was created successfully by checking the number of rows affected    //means we are checking the number of rows affected by the query execution to determine if the user was created successfully. If the number of rows affected is zero, it means that no user was created, and we return an error indicating that the user creation failed. If there are rows affected, it means that the user was created successfully, and we can log this information.


    lastInsertID, errRow := result.LastInsertId()

    if errRow != nil {
        fmt.Println("Error fetching last insert ID:", errRow)
        return nil, errRow
    }

    user := &models.User{
        Id:     lastInsertID, // You can fetch the last inserted ID if needed
        Username: username,
        Email:    email,
    }

    // Log the successful creation of the user
    fmt.Println("User created successfully:", user)

    return user, nil

}


func (u *UserRepositoryImpl) GetByID(id int64) (*models.User, error) {
    // Implement the logic to fetch a user by ID from the database

    query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE id = ?"

    // Execute the query and scan the result into a User struct
    //means we are executing the query with the provided ID and getting a single row as a result. The row variable will hold the result of the query execution, which can be scanned into a User struct to fetch the user details from the database.


    row := u.db.QueryRow(query, id)



    // Initialize an empty User struct to hold the fetched user details
    //means we are creating an empty User struct that will be used to store the user details fetched from the database. This struct will be populated with the data retrieved from the query execution.
    user := &models.User{}

    // Scan the result into the User struct
    //means we are scanning the data from the row into the User struct. The Scan method will populate the fields of the User struct with the corresponding values from the database row. If there is an error during scanning (e.g., if no user is found with the provided ID), it will be handled accordingly.


    err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)


    // Check for errors during scanning
    //means we are checking if there was an error during the scanning process. If the error is sql.ErrNoRows, it means that no user was found with the provided ID, and we return nil without an error. If there is any other error, we log it and return the error to the caller.
    if err != nil {
        if err == sql.ErrNoRows {
            fmt.Println("No user found with ID:", id)
            return nil, nil // No user found, return nil without an error
        }
        fmt.Println("Error fetching user by ID:", err)
        return nil, err // Return the error if it's not a "no rows" error
    }

    fmt.Println("User fetched successfully:", user)

    return user, nil // Return the fetched user and no error
}








func (u *UserRepositoryImpl) GetALL() ([]*models.User, error) {

    // Implement the logic to fetch all users from the database
    query := "SELECT id, username, email, password, created_at, updated_at FROM users"

    // Execute the query and scan the results into a slice of User structs
    // means we are executing the query and getting a set of rows as a result.
    // The rows variable will hold the result of the query execution, which can be iterated over to fetch each user record from the database.
    rows, err := u.db.Query(query)

    // Check for errors during query execution
    if err != nil {
        fmt.Println("Error fetching all users:", err)
        return nil, err // Return the error if the query execution fails
    }

    // Ensure that the rows are closed after processing to prevent resource leaks
    defer rows.Close()

    // Initialize a slice to hold the fetched users
    // means we are creating an empty slice of pointers to User structs.
    // This slice will be used to store the user records fetched from the database.
    var users []*models.User

    // Iterate over the rows and scan each user record into a User struct
    for rows.Next() {
        u := &models.User{} // Create a new User struct for each row

        // Scan the current row into the User struct
        err := rows.Scan(
            &u.Id,
            &u.Username,
            &u.Email,
            &u.Password,
            &u.CreatedAt,
            &u.UpdatedAt,
        )

        if err != nil {
            fmt.Println("Error scanning user record:", err)
            return nil, err // Return the error if scanning fails
        }

        // Append the scanned user to the slice
        users = append(users, u)
    }

    // Check for errors that may have occurred during iteration
    if err = rows.Err(); err != nil {
        fmt.Println("Error iterating over user records:", err)
        return nil, err // Return the error if iteration fails
    }

    fmt.Println("Users fetched successfully:", users)

    return users, nil // Return the slice of users and no error
}



func (u *UserRepositoryImpl) DeleteById(id int64) error {
    // Implement the logic to delete a user by ID from the database
    query := "DELETE FROM users WHERE id = ?"

    // Execute the query with the provided ID
    result, err := u.db.Exec(query, id)

    // Check for errors during query execution
    //means we are checking if there was an error during the execution of the delete query. If there is an error, we log it and return the error to the caller. This allows us to handle any issues that may arise during the user deletion process, such as database connection problems or constraints violations.
    if err != nil {
        fmt.Println("Error deleting user by ID:", err)
        return err // Return the error if the query execution fails
    }

    // Check if the user was deleted successfully by checking the number of rows affected
    //means we are checking the number of rows affected by the delete query to determine if the user was deleted successfully. If the number of rows affected is zero, it means that no user was deleted (possibly because no user with the provided ID exists), and we return an error indicating that the user deletion failed. If there are rows affected, it means that the user was deleted successfully, and we can log this information.
    rowsAffected, errRow := result.RowsAffected()
    if errRow != nil {
        fmt.Println("Error fetching rows affected:", errRow)
        return errRow // Return the error if fetching rows affected fails
    }

    if rowsAffected == 0 {
        fmt.Println("No user found with ID:", id)
        return fmt.Errorf("user not found")
    }

    fmt.Println("User deleted successfully")
    return nil // Return no error if the user was deleted successfully  
}







func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {

    // Implement the logic to fetch a user by email from the database

    query := "SELECT id, username, email, password, created_at, updated_at FROM users WHERE email = ?"

    // Execute the query and scan the result into a User struct
    row := u.db.QueryRow(query, email)

    user := &models.User{} // Create a new User struct

    // Scan the result into the User struct
    err := row.Scan(
        &user.Id,
        &user.Username,
        &user.Email,
        &user.Password,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
    if err != nil {
    if err == sql.ErrNoRows {
        fmt.Println("No user found with email:", email)
        return nil, fmt.Errorf("user not found")
    }
    fmt.Println("Error fetching user by email:", err)
    return nil, err
    }

    return user, nil
}