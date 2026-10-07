package storage

import "github.com/rvxt21/bucket-inventory/pkg/errors"

var (
	ErrDeleteObject  = &errors.Error{Message: "failed to delete file from bucket"}
	ErrUploadObject  = &errors.Error{Message: "failed to upload file to bucket"}
	ErrGetObject     = &errors.Error{Message: "failed to get object from bucket"}
	ErrPresignObject = &errors.Error{Message: "failed to presign bucket object"}
)
