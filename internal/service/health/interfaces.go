package health

import "context"

type Service interface {
	Ready(ctx context.Context) error
}

//nolint:iface // separate types so fx can inject Postgres and S3 separately
type Database interface {
	Ping(ctx context.Context) error
}

//nolint:iface // separate types so fx can inject Postgres and S3 separately
type Storage interface {
	Ping(ctx context.Context) error
}
