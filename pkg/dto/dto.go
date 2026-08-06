package dto

import "io"

type UploadFile struct {
	Filename    string `validate:"required"`
	ContentType string `validate:"required"`
	File        io.Reader
}

type GetFileRequest struct {
	Filename string `param:"name" validate:"required"`
}

type GetFileResponse struct {
	File        []byte
	ContentType *string
}
