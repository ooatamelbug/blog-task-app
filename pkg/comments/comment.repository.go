package comments

import (
	"github.com/ooatamelbug/blog-task-app/pkg/common/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(post models.Comment) (models.Comment, error)
	Update(post models.Comment) (models.Comment, error)
	Delete(post models.Comment) (models.Comment, error)
	FindOne(postId uint64) models.Comment
	Find() []models.Comment
}

type commentRepository struct {
	commentTable *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{
		commentTable: db,
	}
}

func (commenRelation *commentRepository) Create(post models.Comment) (models.Comment, error) {
	err := commenRelation.commentTable.Save(post)
	commenRelation.commentTable.Preload("User").Find(&post)
	return post, err.Error
}

func (commenRelation *commentRepository) Update(post models.Comment) (models.Comment, error) {
	err := commenRelation.commentTable.Save(post)
	commenRelation.commentTable.Preload("User").Find(&post)
	return post, err.Error
}

func (commenRelation *commentRepository) Delete(post models.Comment) (models.Comment, error) {
	err := commenRelation.commentTable.Delete(post)
	return post, err.Error
}

func (commenRelation *commentRepository) FindOne(commentId uint64) models.Comment {
	var comment models.Comment
	commenRelation.commentTable.Preload("User").Preload("Post").Find(&comment, commentId)
	return comment
}

func (commenRelation *commentRepository) Find() []models.Comment {
	var comments []models.Comment
	commenRelation.commentTable.Preload("User").Preload("Post").Find(&comments)
	return comments
}
