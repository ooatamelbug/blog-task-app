package dto

type CreateCommentDto struct {
	Body   string `json:"body" form:"body" binding:"required" validate:"min=3"`
	PostID uint64 `json:"post_id,omitempty" form:"post_id" binding:"required"`
	UserID uint64 `json:"user_id,omitempty" form:"user_id" binding:"required"`
}
