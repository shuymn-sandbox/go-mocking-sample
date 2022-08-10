package persistent

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/shuymn-sandbox/go-mocking-sample/domain/entity"
	"github.com/shuymn-sandbox/go-mocking-sample/domain/repository"
)

const postS3Bucket = "posts"

type postPersistentS3 interface {
	UploadWithContext(aws.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

func NewPostRepository(db *sqlx.DB, s3client postPersistentS3) repository.PostRepository {
	return &postRepositoryImpl{
		db: db,
		s3: s3client,
	}
}

type postRepositoryImpl struct {
	db *sqlx.DB
	s3 postPersistentS3
}

type postMySQL struct {
	ID          int          `db:"id"`
	AuthorID    int          `db:"author_id"`
	Title       string       `db:"title"`
	ContentKey  string       `db:"content_key"`
	PublishedAt sql.NullTime `db:"published_at"`
	CreatedAt   time.Time    `db:"created_at"`
	UpdatedAt   time.Time    `db:"updated_at"`
}

func (p *postRepositoryImpl) CreatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	key := fmt.Sprintf("%d/%s", post.GetAuthorID(), xid.New().String())
	_, err := p.s3.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(postS3Bucket),
		Key:    aws.String(key),
		Body:   post.GetContent(),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	now := time.Now()

	q := `
INSERT INTO posts (author_id, title, content_key, puhlished_at, created_at, updated_at)
VALUES (:author_id, :title, :content_key, :published_at, :created_at, :updated_at)
`
	result, err := p.db.NamedExecContext(ctx, q, postMySQL{
		AuthorID:    post.GetAuthorID(),
		Title:       post.GetTitle(),
		ContentKey:  key,
		PublishedAt: post.GetPublishedAt().NullTime,
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return post.WithID(int(id)).WithCreatedAt(now).WithUpdatedAt(now), nil
}

func (p *postRepositoryImpl) GetPost(ctx context.Context, id int) (*entity.Post, error) {
	panic("not implemented")
}

func (p *postRepositoryImpl) ListPosts(ctx context.Context) ([]*entity.Post, error) {
	panic("not implemented")
}

func (p *postRepositoryImpl) UpdatePost(ctx context.Context, post *entity.Post) (*entity.Post, error) {
	panic("not implemented")
}

func (p *postRepositoryImpl) DeletePost(ctx context.Context, id int) error {
	panic("not implemented")
}
