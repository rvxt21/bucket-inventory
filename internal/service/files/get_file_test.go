package files

import (
	"errors"
	"testing"
	"time"

	"github.com/rvxt21/bucket-inventory/config"
	"github.com/rvxt21/bucket-inventory/internal/database"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
	"github.com/rvxt21/bucket-inventory/test/mocks"
	"github.com/stretchr/testify/require"
)

var (
	linkTTL = 5 * time.Minute

	errDB = errors.New("database error")
	errS3 = errors.New("s3 error")
)

func TestGetFile_Success(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)

	req := &dto.GetFileRequest{
		ID: "123",
	}
	now := time.Now()

	m.db.EXPECT().GetFileByID(ctx, "123").Return(&dto.File{
		ID:          "123",
		Name:        "test",
		StorageKey:  "example",
		ContentType: "application/pdf",
		Size:        1024,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil)

	m.s3.EXPECT().PresignGetObject(ctx, "example", linkTTL).Return("link", nil)

	resp, err := svc.GetFileByID(ctx, req)

	require.NoError(t, err)
	require.Equal(t, &dto.FileResponse{
		ID:          "123",
		Name:        "test",
		ContentType: "application/pdf",
		Size:        1024,
		CreatedAt:   now,
		UpdatedAt:   now,
		Link:        "link",
	}, resp)
}

func TestGetFile_NotFound(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)
	req := &dto.GetFileRequest{ID: "123"}

	m.db.EXPECT().GetFileByID(ctx, "123").Return(nil, database.ErrNotFound)

	resp, err := svc.GetFileByID(ctx, req)

	require.ErrorIs(t, err, ErrFileNotFound)
	require.Nil(t, resp)
}

type MockServices struct {
	db *mocks.FilesDatabase
	s3 *mocks.FilesS3
}

func service(t *testing.T) (*FileService, *MockServices) {
	t.Helper()

	cfg := &config.Config{}
	cfg.LinkTTL = linkTTL
	db := mocks.NewFilesDatabase(t)
	s3 := mocks.NewFilesS3(t)

	return NewFileService(cfg, s3, db), &MockServices{db: db, s3: s3}
}
