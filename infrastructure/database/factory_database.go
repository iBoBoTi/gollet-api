package database

import (
	"errors"
	"github.com/iBoBoTi/gollet-api/infrastructure/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Postgres *pgxpool.Pool
}

var (
	errInvalidDatabaseInstance = errors.New("invalid database instance")
)

const (
	InstancePostgres int = iota
)

// NewDatabaseFactory returns Db type based of the database instance provided
func NewDatabaseFactory(instance int) (*DB, error) {
	switch instance {
	case InstancePostgres:
		return ConnectPostgres(config.NewConfig())
	default:
		return nil, errInvalidDatabaseInstance
	}
}
