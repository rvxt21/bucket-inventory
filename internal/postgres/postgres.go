package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/url"

	"github.com/rvxt21/bucket-inventory/config"

	_ "github.com/lib/pq" // postgres driver
)

type Postgres struct {
	config config.Postgres
	db     *sql.DB
}

func (p *Postgres) Start(ctx context.Context) error {
	dsn := (&url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(p.config.User, p.config.Password),
		Host:     net.JoinHostPort(p.config.Host, p.config.Port),
		Path:     p.config.Database,
		RawQuery: url.Values{"sslmode": {p.config.SSLMode}}.Encode(),
	}).String()

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrOpenPostgresConnection, err)
	}

	db.SetMaxOpenConns(p.config.MaxOpenConns)
	db.SetMaxIdleConns(p.config.MaxIdleConns)
	db.SetConnMaxLifetime(p.config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(p.config.ConnMaxIdleTime)

	if err = db.PingContext(ctx); err != nil {
		db.Close()
		return fmt.Errorf("%w: %w", ErrPingingDatabase, err)
	}

	p.db = db

	return nil
}

func (p *Postgres) Stop(ctx context.Context) error {
	if p.db != nil {
		return p.db.Close()
	}

	return nil
}

func NewPostgres(cfg *config.Config) *Postgres {
	return &Postgres{
		config: cfg.Postgres,
	}
}
