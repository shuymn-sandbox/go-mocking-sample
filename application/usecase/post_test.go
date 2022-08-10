package usecase

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"

	"github.com/shuymn-sandbox/go-mocking-sample/domain/entity"
	"github.com/shuymn-sandbox/go-mocking-sample/domain/repository_mock"
	"github.com/shuymn-sandbox/go-mocking-sample/infrastructure/persistent"
	"github.com/shuymn-sandbox/go-mocking-sample/infrastructure/persistent/mysql"
	"github.com/shuymn-sandbox/go-mocking-sample/infrastructure/persistent/s3"
	"github.com/shuymn-sandbox/go-mocking-sample/infrastructure/persistent/s3_mock"
)

func TestPostUsecase_CreatePost_1(t *testing.T) {
	mockPostRepository := &repository_mock.PostRepository{}
	mockPostRepository.On("CreatePost", mock.Anything, mock.Anything).Return(entity.NewPost().WithID(1), nil)

	uc := &postUsecaseImpl{
		postRepository: mockPostRepository,
	}
	post, err := uc.CreatePost(context.Background(), entity.NewUser(), &CreatePostInput{})
	if err != nil {
		t.Fatal(err)
	}
	if got := post.GetID(); got != 1 {
		t.Errorf("id must be 1, but %d", got)
	}
}

func TestPostUsecase_CreatePost_TheOther(t *testing.T) {
	db, err := sqlx.Open("mysql", "")
	if err != nil {
		t.Fatal(err)
	}

	mockS3 := &s3_mock.PostS3Persistent{}
	mockS3.On("UploadContent", mock.Anything, mock.Anything).Return(&s3.UploadContentOutput{Key: "test"}, nil)

	uc := &theOtherPostUsecaseImpl{
		s3:    mockS3,
		mysql: mysql.NewPostMySQLPersistent(db),
	}
	post, err := uc.CreatePost(context.Background(), entity.NewUser(), &CreatePostInput{})
	if err != nil {
		t.Fatal(err)
	}
	if got := post.GetID(); got != 1 {
		t.Errorf("id must be 1, but %d", got)
	}
}

type mockPostPersistentS3Impl struct {
	uploadWithContextFunc func(aws.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

func (m *mockPostPersistentS3Impl) UploadWithContext(ctx aws.Context, input *s3manager.UploadInput, opts ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	return m.uploadWithContextFunc(ctx, input, opts...)
}

func TestPostUsecase_CreatePost_2(t *testing.T) {
	db, err := sqlx.Open("mysql", "")
	if err != nil {
		t.Fatal(err)
	}

	mockS3 := &mockPostPersistentS3Impl{
		uploadWithContextFunc: func(_ aws.Context, _ *s3manager.UploadInput, _ ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
			return nil, nil
		},
	}

	uc := &postUsecaseImpl{
		postRepository: persistent.NewPostRepository(db, mockS3),
	}
	post, err := uc.CreatePost(context.Background(), entity.NewUser(), &CreatePostInput{})
	if err != nil {
		t.Fatal(err)
	}
	if got := post.GetID(); got != 1 {
		t.Errorf("id must be 1, but %d", got)
	}
}
