package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	Dbname   string
	User     string
	Password string
}

type Database interface {
	NewConnection() (*pgxpool.Pool, error)
}

func NewDatabase(databse string, dbConfig DBConfig) Database {
	switch databse {
	case "postgres":
		return postgresDatabase{dbConfig}
	default:
		return nil
	}
}

type postgresDatabase struct {
	config DBConfig
}

func (pg postgresDatabase) NewConnection() (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
		pg.config.Host,
		pg.config.Port,
		pg.config.User,
		pg.config.Password,
		pg.config.Dbname,
	)

	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("error to connect to database: %v", err)
	}

	return conn, nil
}
