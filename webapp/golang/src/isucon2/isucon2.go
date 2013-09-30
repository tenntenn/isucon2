package isucon2

import (
	"database/sql"
	"log"
)

var (
	Conf *Config
	Db   *sql.DB
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
