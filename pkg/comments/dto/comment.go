package dto

type CreateCommentDto struct {
	Body string `json:"body" form:"body" binding:"required" validate:"min:3"`
	User uint64 `json:"user" form:"user" binding:"required"`
	Post uint64 `json:"post" form:"post" binding:"required"`
}
