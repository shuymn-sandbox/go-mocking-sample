package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
	"github.com/rs/xid"
)

const s3PostS3Bucket = "posts"

//go:generate mockery --dir . --name PostS3Persistent --outpkg s3_mock --output ../s3_mock --case underscore
type PostS3Persistent interface {
	UploadContent(ctx context.Context, input *UploadContentInput) (*UploadContentOutput, error)
}

type postS3PersistentS3 interface {
	UploadWithContext(aws.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

func NewPostS3Persistent(s3client postS3PersistentS3) PostS3Persistent {
	return &postS3PersistentImpl{
		s3: s3client,
	}
}

type postS3PersistentImpl struct {
	s3 postS3PersistentS3
}

type UploadContentInput struct {
	UserID  int
	Content io.Reader
}

type UploadContentOutput struct {
	Key string
}

func (p *postS3PersistentImpl) UploadContent(ctx context.Context, input *UploadContentInput) (*UploadContentOutput, error) {
	key := fmt.Sprintf("%d/%s", input.UserID, xid.New().String())
	_, err := p.s3.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(s3PostS3Bucket),
		Key:    aws.String(key),
		Body:   input.Content,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &UploadContentOutput{
		Key: key,
	}, nil
}
