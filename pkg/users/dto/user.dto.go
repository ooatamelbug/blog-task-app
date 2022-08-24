package dto

type UserCreateDTO struct {
	FirstName string `json:"firstname" form:"firstname" binding:"required" validate:"min:3"`
	LastName  string `json:"lastname" form:"lastname" binding:"required" validate:"min:3"`
	Email     string `json:"email" form:"email" binding:"required" validate:"min:3"`
	Password  string `json:"password" form:"password" binding:"required" validate:"min:3"`
}
