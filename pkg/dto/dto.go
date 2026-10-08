package dto

import (
	"io"
	"time"
)

type CreateFile struct {
	Name        string
	StorageKey  string
	ContentType string
	Size        int64
}

type File struct {
	ID          string
	Name        string
	StorageKey  string
	ContentType string
	Size        int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UploadFile struct {
	Filename    string `validate:"required"`
	ContentType string `validate:"required"`
	Size        int64
	File        io.Reader
}

type GetFileRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type DeleteFileRequest struct {
	ID string `param:"id" validate:"required,uuid"`
}

type GetFileResponse struct {
	File        []byte
	ContentType *string
}

type FileResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	ContentType string    `json:"content_type,omitempty"`
	Size        int64     `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Link        string    `json:"link,omitempty"`
}
