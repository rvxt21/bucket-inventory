package health

import "github.com/rvxt21/bucket-inventory/pkg/errors"

var (
	ErrPingDatabase = &errors.Error{Message: "error ping database"}
	ErrPingS3       = &errors.Error{Message: "error ping s3 service"}
)
