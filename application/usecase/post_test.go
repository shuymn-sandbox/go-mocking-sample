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

// Repositoryを分離したパターン
func TestPostUsecase_CreatePost_2(t *testing.T) {
	// DBはmockしない
	db, err := sqlx.Open("mysql", "")
	if err != nil {
		t.Fatal(err)
	}

	// S3はmockeryで生成されたmockを使う
	mockS3 := &s3_mock.PostS3Persistent{}
	mockS3.On("UploadContent", mock.Anything, mock.Anything).Return(&s3.UploadContentOutput{Key: "test"}, nil)

	// usecaseを初期化
	uc := &theOtherPostUsecaseImpl{
		s3:    mockS3,
		mysql: mysql.NewPostMySQLPersistent(db),
	}

	// CreatePostを呼ぶ
	post, err := uc.CreatePost(context.Background(), entity.NewUser(), &CreatePostInput{})
	if err != nil {
		t.Fatal(err)
	}
	// 返り値をテストする
	if got := post.GetID(); got != 1 {
		t.Errorf("id must be 1, but %d", got)
	}
}

// mockeryで生成されたPostRepositoryのmockを使う
func TestPostUsecase_CreatePost_1_問題あり(t *testing.T) {
	// mockを初期化
	mockPostRepository := &repository_mock.PostRepository{}
	mockPostRepository.On("CreatePost", mock.Anything, mock.Anything).Return(entity.NewPost().WithID(1), nil)

	// usecaseを初期化
	uc := &postUsecaseImpl{
		postRepository: mockPostRepository,
	}

	// CreatePostを呼ぶ
	post, err := uc.CreatePost(context.Background(), entity.NewUser(), &CreatePostInput{})
	if err != nil {
		t.Fatal(err)
	}
	// 返り値をテストする
	if got := post.GetID(); got != 1 {
		t.Errorf("id must be 1, but %d", got)
	}
}

// Repositoryを分離しないパターン
// mockeryのmockではなく自前でmockを用意するパターン
func TestPostUsecase_CreatePost_1_解決策(t *testing.T) {
	// dbはmockしない
	db, err := sqlx.Open("mysql", "")
	if err != nil {
		t.Fatal(err)
	}

	// s3は自前で用意したmockを使う
	mockS3 := &mockPostPersistentS3Impl{
		uploadWithContextFunc: func(_ aws.Context, _ *s3manager.UploadInput, _ ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
			// なにもしない
			return &s3manager.UploadOutput{}, nil
		},
	}

	// usecaseを初期化
	uc := &postUsecaseImpl{
		postRepository: persistent.NewPostRepository(db, mockS3),
	}

	// CreatePostを呼ぶ
	post, err := uc.CreatePost(context.Background(), entity.NewUser(), &CreatePostInput{})
	if err != nil {
		t.Fatal(err)
	}
	// 返り値をテストする
	if got := post.GetID(); got != 1 {
		t.Errorf("id must be 1, but %d", got)
	}
}

// persistent.postPersistentS3のmock
type mockPostPersistentS3Impl struct {
	uploadWithContextFunc func(aws.Context, *s3manager.UploadInput, ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

func (m *mockPostPersistentS3Impl) UploadWithContext(ctx aws.Context, input *s3manager.UploadInput, opts ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error) {
	if m == nil {
		return nil, nil
	}
	return m.uploadWithContextFunc(ctx, input, opts...)
}
