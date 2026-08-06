package service

import "errors"

var (
	ErrUploadFile = errors.New("error upload file")
	ErrGetFile    = errors.New("error get file")
)
