package files

import (
	"testing"
	"time"

	"github.com/rvxt21/bucket-inventory/pkg/dto"
	"github.com/stretchr/testify/require"
)

func TestGetFiles_Success(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)

	now := time.Now()
	m.db.EXPECT().List(ctx).Return([]dto.File{
		{
			ID:          "123",
			Name:        "a.pdf",
			StorageKey:  "123-key",
			ContentType: "application/pdf",
			Size:        1024,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
		{ID: "456", StorageKey: "456-key"},
	}, nil).Once()

	m.s3.EXPECT().PresignGetObject(ctx, "123-key", linkTTL).Return("link-123", nil).Once()
	m.s3.EXPECT().PresignGetObject(ctx, "456-key", linkTTL).Return("link-456", nil).Once()

	resp, err := svc.GetFiles(ctx)
	require.NoError(t, err)
	require.Equal(t, []dto.FileResponse{
		{
			ID:          "123",
			Name:        "a.pdf",
			ContentType: "application/pdf",
			Size:        1024,
			CreatedAt:   now,
			UpdatedAt:   now,
			Link:        "link-123",
		},
		{ID: "456", Link: "link-456"},
	}, resp)
}

func TestGetFiles_Empty(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)

	m.db.EXPECT().List(ctx).Return([]dto.File{}, nil).Once()

	resp, err := svc.GetFiles(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Empty(t, resp)
}

func TestGetFiles_DBError(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)

	m.db.EXPECT().List(ctx).Return(nil, errDB).Once()

	resp, err := svc.GetFiles(ctx)
	require.ErrorIs(t, err, errDB)
	require.Nil(t, resp)
}

func TestGetFiles_PresignError(t *testing.T) {
	ctx := t.Context()

	svc, m := service(t)

	m.db.EXPECT().List(ctx).Return([]dto.File{
		{ID: "1", StorageKey: "key-1"},
		{ID: "2", StorageKey: "key-2"},
		{ID: "3", StorageKey: "key-3"},
	}, nil).Once()

	m.s3.EXPECT().PresignGetObject(ctx, "key-1", linkTTL).Return("link-1", nil).Once()
	m.s3.EXPECT().PresignGetObject(ctx, "key-2", linkTTL).Return("", errS3).Once()

	resp, err := svc.GetFiles(ctx)
	require.ErrorIs(t, err, errS3)
	require.Nil(t, resp)
}
