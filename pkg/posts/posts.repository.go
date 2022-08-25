package posts

import (
	"github.com/ooatamelbug/blog-task-app/pkg/common/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post models.Post) (models.Post, error)
	Update(post models.Post) (models.Post, error)
	Delete(post models.Post) (models.Post, error)
	FindOne(postId uint64) models.Post
	Find() []models.Post
}

type postRepository struct {
	postTable *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{
		postTable: db,
	}
}

func (postRelation *postRepository) Create(post models.Post) (models.Post, error) {
	err := postRelation.postTable.Save(&post)
	if err != nil {
		return post, err.Error
	}
	postRelation.postTable.Preload("User").Find(&post)
	return post, nil
}

func (postRelation *postRepository) Update(post models.Post) (models.Post, error) {
	err := postRelation.postTable.Save(&post)
	postRelation.postTable.Preload("User").Find(&post)
	return post, err.Error
}

func (postRelation *postRepository) Delete(post models.Post) (models.Post, error) {
	err := postRelation.postTable.Delete(&post)
	return post, err.Error
}

func (postRelation *postRepository) FindOne(postId uint64) models.Post {
	var post models.Post
	postRelation.postTable.Preload("User").Preload("CommentsList").Find(&post, postId)
	return post
}

func (postRelation *postRepository) Find() []models.Post {
	var posts []models.Post
	postRelation.postTable.Preload("User").Preload("CommentsList").Find(&posts)
	return posts
}
