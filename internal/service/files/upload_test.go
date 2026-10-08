package files

import (
	"strings"
	"testing"

	"github.com/rvxt21/bucket-inventory/pkg/dto"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func uploadRequest() dto.UploadFile {
	return dto.UploadFile{
		Filename:    "report.pdf",
		ContentType: "application/pdf",
		Size:        1024,
		File:        strings.NewReader("content"),
	}
}

func TestUploadFile_Success(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)
	req := uploadRequest()

	created := &dto.File{ID: "123", Name: req.Filename}

	m.s3.EXPECT().UploadObject(ctx, mock.Anything, req.File, req.ContentType).Return(nil).Once()
	m.db.EXPECT().Create(ctx, mock.Anything).Return(created, nil).Once()

	resp, err := svc.UploadFile(ctx, req)
	require.NoError(t, err)
	require.Equal(t, &dto.FileResponse{ID: "123", Name: req.Filename}, resp)
}

func TestUploadFile_S3Error(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)
	req := uploadRequest()

	m.s3.EXPECT().UploadObject(ctx, mock.Anything, req.File, req.ContentType).Return(errS3).Once()

	resp, err := svc.UploadFile(ctx, req)
	require.ErrorIs(t, err, errS3)
	require.Nil(t, resp)
}

func TestUploadFile_DBError(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)
	req := uploadRequest()

	m.s3.EXPECT().UploadObject(ctx, mock.Anything, req.File, req.ContentType).Return(nil).Once()

	m.db.EXPECT().Create(ctx, mock.Anything).Return(nil, errDB).Once()

	resp, err := svc.UploadFile(ctx, req)
	require.ErrorIs(t, err, errDB)
	require.Nil(t, resp)
}
