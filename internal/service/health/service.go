package health

import "context"

type Health struct {
	postgres Database
	s3       Storage
}

func (h *Health) Ready(ctx context.Context) error {
	if err := h.postgres.Ping(ctx); err != nil {
		return ErrPingDatabase.Wrap(err)
	}

	if err := h.s3.Ping(ctx); err != nil {
		return ErrPingS3.Wrap(err)
	}

	return nil
}

func NewHealthService(postgres Database, s3 Storage) *Health {
	return &Health{
		postgres: postgres,
		s3:       s3,
	}
}
