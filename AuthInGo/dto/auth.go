package dto


//


type LoginUserRequestDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type CreateUserRequestDTO struct {

	// The `json:"username"` tag specifies that when this struct is serialized to JSON, the field will be represented as "username". The `validate:"required,min=3,max=50"` tag indicates that this field is required and must have a minimum length of 3 characters and a maximum length of 50 characters.

	Username string `json:"username" validate:"required,min=3,max=50"`

	// The `json:"email"` tag specifies that when this struct is serialized to JSON, the field will be represented as "email". The `validate:"required,email"` tag indicates that this field is required and must be a valid email address.

	Email    string `json:"email" validate:"required,email"`


	// The `json:"password"` tag specifies that when this struct is serialized to JSON, the field will be represented as "password". The `validate:"required,min=8"` tag indicates that this field is required and must have a minimum length of 8 characters.

	Password string `json:"password" validate:"required,min=8"`
}

// {"email", "password"}


