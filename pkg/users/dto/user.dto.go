package dto

type CreateUserDTO struct {
	FirstName string `json:"firstname" form:"firstname" binding:"required" validate:"min:3"`
	LastName  string `json:"lastname" form:"lastname" binding:"required" validate:"min:3"`
	Email     string `json:"email" form:"email" binding:"required" validate:"min:3"`
	Password  string `json:"password" form:"password" binding:"required" validate:"min:3"`
}

type SearchUser struct {
	Email  string `json:"email" form:"email" `
	UserId string `json:"userid" form:"userid" `
}

type UpdateUserDTO struct {
	FirstName string `json:"firstname" form:"firstname"`
	LastName  string `json:"lastname" form:"lastname"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	ID        uint64 `json:"id" form:"id" binding:"required"`
}

type SearchWithAnd struct {
	Email    string `json:"email" form:"email" `
	Password string `json:"password" form:"password"`
	ID       uint64 `json:"id" form:"id" binding:"required"`
}
