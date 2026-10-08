package database

import (
	"github.com/rvxt21/bucket-inventory/pkg/errors"
)

var (
	ErrOpenPostgresConnection = &errors.Error{Message: "error opening postgres connection"}
	ErrPingingDatabase        = &errors.Error{Message: "error pinging database"}
	ErrNotFound               = &errors.Error{Message: "requested data not found"}
	ErrCreateFile             = &errors.Error{Message: "error creating file"}
	ErrGetFileByID            = &errors.Error{Message: "error get file by id"}
	ErrListFiles              = &errors.Error{Message: "error listing files"}
	ErrDeleteFile             = &errors.Error{Message: "error deleting file from database"}
)
