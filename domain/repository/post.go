package repository

import (
	"context"

	"github.com/shuymn-sandbox/go-mocking-sample/domain/entity"
)

//go:generate mockery --dir . --name PostRepository --outpkg repository_mock --output ../repository_mock --case underscore
type PostRepository interface {
	CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error)
	GetPost(ctx context.Context, id int) (*entity.Post, error)
	ListPosts(ctx context.Context) ([]*entity.Post, error)
	UpdatePost(ctx context.Context, post *entity.Post) (*entity.Post, error)
	DeletePost(ctx context.Context, id int) error
}
