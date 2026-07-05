package dto

import "io"

type UploadFile struct {
	Filename    string
	ContentType string
	File        io.Reader
}
