package posts

import (
	"errors"

	"github.com/mashingan/smapping"
	"github.com/ooatamelbug/blog-task-app/pkg/common/models"
	"github.com/ooatamelbug/blog-task-app/pkg/posts/dto"
)

type PostService interface {
	CreatePost(post dto.CreatePostDto) (models.Post, error)
	UpdatePost(post dto.CreatePostDto, postId uint64) (models.Post, error)
	DeletePost(postId uint64, userId uint64) (models.Post, error)
	GetPost(postId uint64) (models.Post, error)
	GetPosts() []models.Post
}

type postService struct {
	postRepository PostRepository
}

func NewPostService(postRepo PostRepository) PostService {
	return &postService{
		postRepository: postRepo,
	}
}

func (postserv *postService) CreatePost(post dto.CreatePostDto) (models.Post, error) {
	newPost := models.Post{}
	err := smapping.FillStruct(&newPost, smapping.MapFields(&post))
	if err != nil {
		return newPost, err
	}
	row, err := postserv.postRepository.Create(newPost)
	if err != nil {
		return newPost, err
	}
	return row, nil
}

func (postserv *postService) UpdatePost(post dto.CreatePostDto, postId uint64) (models.Post, error) {
	getpost := postserv.postRepository.FindOne(postId)
	if post.Title == "" {
		return getpost, errors.New("no post")
	}
	if getpost.User.ID != post.User {
		return getpost, errors.New("not allowed to Update this post")
	}

	updatePost := models.Post{}
	err := smapping.FillStruct(&updatePost, smapping.MapFields(&post))
	if err != nil {
		return updatePost, err
	}
	row, err := postserv.postRepository.Update(updatePost)
	if err != nil {
		return updatePost, err
	}
	return row, nil
}

func (postserv *postService) DeletePost(postId uint64, userId uint64) (models.Post, error) {
	post := postserv.postRepository.FindOne(postId)
	if post.Title == "" {
		return post, errors.New("no post")
	}
	if post.User.ID != userId {
		return post, errors.New("not allowed to delete this post")
	}
	row, err := postserv.postRepository.Delete(post)
	if err != nil {
		return post, err
	}
	return row, nil
}

func (postserv *postService) GetPosts() []models.Post {
	var posts []models.Post
	allPosts := postserv.postRepository.Find()
	posts = allPosts
	return posts
}

func (postserv *postService) GetPost(postId uint64) (models.Post, error) {
	post := postserv.postRepository.FindOne(postId)
	if post.Title == "" {
		return post, errors.New("not found")
	}
	return post, nil
}
