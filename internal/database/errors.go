package database

import "github.com/rvxt21/bucket-inventory/pkg/errors"

var (
	ErrOpenPostgresConnection = &errors.Error{Message: "error opening postgres connection"}
	ErrPingingDatabase        = &errors.Error{Message: "error pinging database"}
)
