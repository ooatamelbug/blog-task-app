package dto

type CreatePostDto struct {
	Title string `json:"title" form:"title" binding:"required" validate:"min:3"`
	Body  string `json:"body" form:"body" binding:"required" validate:"min:3"`
	User  uint64 `json:"user" form:"user" binding:"required"`
}
