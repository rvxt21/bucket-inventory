package files

import (
	"testing"

	"github.com/rvxt21/bucket-inventory/internal/database"
	"github.com/rvxt21/bucket-inventory/pkg/dto"
	"github.com/stretchr/testify/require"
)

func TestDeleteFile_Success(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)
	req := &dto.DeleteFileRequest{ID: "123"}

	m.db.EXPECT().DeleteFile(ctx, "123").Return("key", nil).Once()
	m.s3.EXPECT().DeleteObject(ctx, "key").Return(nil).Once()

	err := svc.DeleteFile(ctx, req)
	require.NoError(t, err)
}

func TestDeleteFile_NotFound(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)
	req := &dto.DeleteFileRequest{ID: "123"}

	m.db.EXPECT().DeleteFile(ctx, "123").Return("", database.ErrNotFound).Once()

	err := svc.DeleteFile(ctx, req)
	require.ErrorIs(t, err, ErrFileNotFound)
}

func TestDeleteFile_DBError(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)
	req := &dto.DeleteFileRequest{ID: "123"}

	m.db.EXPECT().DeleteFile(ctx, "123").Return("", errDB).Once()

	err := svc.DeleteFile(ctx, req)
	require.ErrorIs(t, err, ErrDeleteFile)
	require.ErrorIs(t, err, errDB)
}

func TestDeleteFile_S3Error(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)
	req := &dto.DeleteFileRequest{ID: "123"}

	m.db.EXPECT().DeleteFile(ctx, "123").Return("key", nil).Once()
	m.s3.EXPECT().DeleteObject(ctx, "key").Return(errS3).Once()

	err := svc.DeleteFile(ctx, req)
	require.NoError(t, err)
}
