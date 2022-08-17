package usecase

import (
	"bytes"
	"context"

	"github.com/pkg/errors"
	"github.com/shuymn-sandbox/go-mocking-sample/domain/entity"
	"github.com/shuymn-sandbox/go-mocking-sample/infrastructure/persistent/mysql"
	"github.com/shuymn-sandbox/go-mocking-sample/infrastructure/persistent/s3"
)

// 1 Repository(Persistent) 1外部ソース のバージョン
type theOtherPostUsecaseImpl struct {
	s3    s3.PostS3Persistent
	mysql mysql.PostMySQLPersistent
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
