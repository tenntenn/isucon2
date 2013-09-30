package isucon2

import (
	"fmt"
)

type DbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	UserName string `json:"username"`
	Password string `json:"password"`
	DbName   string `json:"dbname"`
}

func (db *DbConfig) String() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		db.UserName,
		db.Password,
		db.Host,
		db.Port,
		db.DbName,
	)
}
