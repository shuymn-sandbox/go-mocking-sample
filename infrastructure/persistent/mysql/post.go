package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/shuymn-sandbox/go-mocking-sample/domain/entity"
)

//go:generate mockery --dir . --name PostMySQLPersistent --outpkg mysql_mock --output ../mysql_mock --case underscore
type PostMySQLPersistent interface {
	CreatePost(ctx context.Context, input *CreatePostInput) (*entity.Post, error)
}

func NewPostMySQLPersistent(db *sqlx.DB) PostMySQLPersistent {
	return &postMySQLPersistentImpl{
		db: db,
	}
}

type postMySQLPersistentImpl struct {
	db *sqlx.DB
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

type CreatePostInput struct {
	ContentKey string
	Post       *entity.Post
}

func (p *postMySQLPersistentImpl) CreatePost(ctx context.Context, input *CreatePostInput) (*entity.Post, error) {
	now := time.Now()

	q := `
INSERT INTO posts (author_id, title, content_key, puhlished_at, created_at, updated_at)
VALUES (:author_id, :title, :content_key, :published_at, :created_at, :updated_at)
`
	post := input.Post
	result, err := p.db.NamedExecContext(ctx, q, postMySQL{
		AuthorID:    post.GetAuthorID(),
		Title:       post.GetTitle(),
		ContentKey:  input.ContentKey,
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
