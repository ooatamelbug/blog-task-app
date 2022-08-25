package comments

import (
	"errors"
	"log"

	"github.com/mashingan/smapping"
	"github.com/ooatamelbug/blog-task-app/pkg/comments/dto"
	"github.com/ooatamelbug/blog-task-app/pkg/common/models"
)

type CommentService interface {
	CreateComment(comment dto.CreateCommentDto) (models.Comment, error)
	UpdateComment(comment dto.CreateCommentDto, postId uint64) (models.Comment, error)
	DeleteComment(commentId uint64, userId uint64) (models.Comment, error)
	GetComment(commentId uint64) (models.Comment, error)
	GetComments() []models.Comment
}

type commentService struct {
	commentRepository CommentRepository
}

func NewCommentService(commentRepo CommentRepository) CommentService {
	return &commentService{
		commentRepository: commentRepo,
	}
}

func (commentServ *commentService) CreateComment(comment dto.CreateCommentDto) (models.Comment, error) {
	newComment := models.Comment{}
	err := smapping.FillStruct(&newComment, smapping.MapFields(&comment))
	if err != nil {
		return newComment, err
	}
	row, err := commentServ.commentRepository.Create(newComment)
	if err != nil {
		return newComment, err
	}
	return row, nil
}

func (commentServ *commentService) UpdateComment(comment dto.CreateCommentDto, commentId uint64) (models.Comment, error) {
	getcomment := commentServ.commentRepository.FindOne(commentId)
	log.Println(getcomment)
	if comment.Body == "" {
		return getcomment, errors.New("no comment")
	}

	if getcomment.User.ID != comment.UserID {
		return getcomment, errors.New("not allowed to Update this comment")
	}

	updateComment := models.Comment{}
	err := smapping.FillStruct(&updateComment, smapping.MapFields(&comment))
	if err != nil {
		return updateComment, err
	}
	updateComment.ID = commentId
	row, err := commentServ.commentRepository.Update(updateComment)
	if err != nil {
		return updateComment, err
	}
	return row, nil
}

func (commentServ *commentService) DeleteComment(commentId uint64, userId uint64) (models.Comment, error) {
	comment := commentServ.commentRepository.FindOne(commentId)
	if comment.Body == "" {
		return comment, errors.New("no comment")
	}
	if comment.User.ID != userId {
		return comment, errors.New("not allowed to delete this comment")
	}
	row, err := commentServ.commentRepository.Delete(comment)
	if err != nil {
		return comment, err
	}
	return row, nil
}

func (commentServ *commentService) GetComments() []models.Comment {
	var comments []models.Comment
	allComments := commentServ.commentRepository.Find()
	comments = allComments
	return comments
}

func (commentServ *commentService) GetComment(commentId uint64) (models.Comment, error) {
	comment := commentServ.commentRepository.FindOne(commentId)
	if comment.Body == "" {
		return comment, errors.New("not found")
	}
	return comment, nil
}
