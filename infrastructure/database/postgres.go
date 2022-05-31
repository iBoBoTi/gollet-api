package database

import (
	"context"
	"fmt"
	"github.com/iBoBoTi/gollet-api/infrastructure/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

// ConnectPostgres connects to the postgres database pool and assigns it the Db struct pool field returning Db
func ConnectPostgres(c *config.Config) (*DB, error) {
	log.Println("Connecting to Postgresql DB pool")
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.DBHost,
		c.DBPort,
		c.DBUser,
		c.DBName,
		c.DBPassword,
	)
	conf, err := pgxpool.ParseConfig(dns)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		return nil, err
	}

	return &DB{Postgres: pool}, nil
}
