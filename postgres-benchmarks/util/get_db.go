package util

import (
	"database/sql"

	"github.com/globalsign/mgo"

	_ "github.com/lib/pq" // Postgres SQL Driver
)

// DBOptions connection configurable options
type DBOptions struct {
	MaxIdleConns int // Default 2
	MaxOpenConns int // Default 30
}

// Default return a new DBOptions instance with default values applied
func (opts DBOptions) Default() DBOptions {
	if opts.MaxIdleConns == 0 {
		opts.MaxIdleConns = 2
	}

	if opts.MaxOpenConns == 0 {
		opts.MaxOpenConns = 30
	}

	return opts
}

// Apply return a new DBOptions instance with values from other instances applied sequentially
func (opts DBOptions) Apply(options ...DBOptions) DBOptions {
	for i := 0; i < len(options); i++ {
		opts.MaxIdleConns = options[i].MaxIdleConns
		opts.MaxOpenConns = options[i].MaxOpenConns
	}

	return opts
}

// GetSQLDatabase return a SQL database connection
func GetSQLDatabase(driver string, connStr string, options ...DBOptions) (*sql.DB, error) {
	db, err := sql.Open(driver, connStr)
	if err != nil {
		return nil, err
	}

	opts := DBOptions{}.Apply(options...).Default()
	db.SetMaxIdleConns(opts.MaxIdleConns)
	db.SetMaxOpenConns(opts.MaxOpenConns)

	return db, nil
}

// GetMongoDatabase return a Mongo database connection
func GetMongoDatabase(connStr string, options ...DBOptions) (*mgo.Session, error) {
	db, err := mgo.Dial(connStr)
	if err != nil {
		return nil, err
	}

	opts := DBOptions{}.Apply(options...).Default()
	db.SetPoolLimit(opts.MaxOpenConns)

	return db, nil
}
