package dbtools

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type DBInit struct {
	db *sql.DB
}

func Connect(dn string, dsn string) (*DBInit, error) {
	dbconn, err := sql.Open(dn, dsn)

	if err != nil {
		log.Fatal(err.Error())
	}

	return &DBInit{
		db: dbconn,
	}, nil
}
