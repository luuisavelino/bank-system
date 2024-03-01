package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	Dbname   string
	User     string
	Password string
}

type Database interface {
	NewConnection() (*gorm.DB, error)
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

func (pg postgresDatabase) NewConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Sao_Paulo",
		pg.config.Host,
		pg.config.Port,
		pg.config.User,
		pg.config.Password,
		pg.config.Dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error to connect to database: %v", err)
	}

	return db, nil
}
