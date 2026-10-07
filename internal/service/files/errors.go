package files

import "github.com/rvxt21/bucket-inventory/pkg/errors"

var (
	ErrUploadFile   = &errors.Error{Message: "error upload file"}
	ErrGetFile      = &errors.Error{Message: "error get file"}
	ErrGetFiles     = &errors.Error{Message: "error get files"}
	ErrFileNotFound = &errors.Error{Message: "requested file not found"}
)
