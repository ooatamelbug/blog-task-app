package authdto

type Login struct {
	Email    string `json:"email" form:"email" binding:"required" validate:"min:3"`
	Password string `json:"password" form:"password" binding:"required" validate:"min:3"`
}
