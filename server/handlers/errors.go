package handlers

import "github.com/rvxt21/bucket-inventory/pkg/errors"

var (
	ErrBind     = &errors.Error{Message: "Bind error"}
	ErrValidate = &errors.Error{Message: "Validation error"}
)
