package usecase

import (
	"bytes"
	"context"

	"github.com/pkg/errors"
	"github.com/shuymn-sandbox/go-mocking-sample/domain/entity"
	"github.com/shuymn-sandbox/go-mocking-sample/domain/repository"
	"gopkg.in/guregu/null.v4"
)

// インターフェース
type PostUsecase interface {
	CreatePost(ctx context.Context, user *entity.User, input *CreatePostInput) (*entity.Post, error)
}

// 1 Repository 複数外部ソース のバージョン
type postUsecaseImpl struct {
	postRepository repository.PostRepository
}

type CreatePostInput struct {
	Title       string
	Content     []byte
	PublishedAt null.Time
}

func (p *postUsecaseImpl) CreatePost(ctx context.Context, user *entity.User, input *CreatePostInput) (*entity.Post, error) {
	// 実際はバリデーションしてから
	post := entity.NewPost().
		WithAuthorID(user.GetID()).
		WithAuthor(user).
		WithTitle(input.Title).
		WithContent(bytes.NewBuffer(input.Content)).
		WithPublishedAt(input.PublishedAt)

	var err error
	post, err = p.postRepository.CreatePost(ctx, post)
	if err != nil {
		// 実際は4XX系のエラーのためにいくつか条件分岐がある
		return nil, errors.WithStack(err)
	}
	return post, nil
}
