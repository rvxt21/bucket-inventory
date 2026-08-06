package postgres

import "errors"

var (
	ErrOpenPostgresConnection = errors.New("error opening postgres connection")
	ErrPingingDatabase        = errors.New("error pinging database")
)
