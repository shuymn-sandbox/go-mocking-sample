package usecase

import (
	"bytes"
	"context"

	"github.com/pkg/errors"
	"github.com/shuymn-sandbox/go-mocking-sample/domain/entity"
	"github.com/shuymn-sandbox/go-mocking-sample/domain/repository"
	"github.com/shuymn-sandbox/go-mocking-sample/infrastructure/persistent/mysql"
	"github.com/shuymn-sandbox/go-mocking-sample/infrastructure/persistent/s3"
	"gopkg.in/guregu/null.v4"
)

type PostUsecase interface {
	CreatePost(ctx context.Context, user *entity.User, input *CreatePostInput) (*entity.Post, error)
}

type postUsecaseImpl struct {
	postRepository repository.PostRepository
}

type theOtherPostUsecaseImpl struct {
	s3    s3.PostS3Persistent
	mysql mysql.PostMySQLPersistent
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

func (p *theOtherPostUsecaseImpl) CreatePost(ctx context.Context, user *entity.User, input *CreatePostInput) (*entity.Post, error) {
	content := bytes.NewBuffer(input.Content)
	// 実際はバリデーションしてから
	o, err := p.s3.UploadContent(ctx, &s3.UploadContentInput{
		UserID:  user.GetID(),
		Content: content,
	})
	if err != nil {
		// 実際は4XX系のエラーのためにいくつか条件分岐がある
		return nil, errors.WithStack(err)
	}
	post, err := p.mysql.CreatePost(ctx, &mysql.CreatePostInput{
		ContentKey: o.Key,
		Post: entity.NewPost().
			WithAuthorID(user.GetID()).
			WithAuthor(user).
			WithTitle(input.Title).
			WithContent(content).
			WithPublishedAt(input.PublishedAt),
	})
	if err != nil {
		// 実際は4XX系のエラーのためにいくつか条件分岐がある
		return nil, errors.WithStack(err)
	}
	return post, nil
}
