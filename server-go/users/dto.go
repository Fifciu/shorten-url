package users

type UpdateUserPasswordDto struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required"`
}

type UpdateEmailPasswordDto struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewEmail        string `json:"new_email" validate:"required,email"`
}
