package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	connStr = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Database struct {
	db *sqlx.DB
}

func NewDatabase(conf DatabaseConfig) (*Database, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf(connStr, conf.Host, conf.Port, conf.User, conf.Password, conf.DBName, conf.SSLMode))
	log.Println(fmt.Sprintf(connStr, conf.Host, conf.Port, conf.User, conf.Password, conf.DBName, conf.SSLMode))
	if err != nil {
		log.Println("!!! ", err)
		return nil, err
	}
	return &Database{db}, nil
}

func (d *Database) GetDB() *sqlx.DB {
	return d.db
}

func (d *Database) CloseDB() error {
	err := d.db.Close()
	if err != nil {
		return err
	}
	return nil
}
