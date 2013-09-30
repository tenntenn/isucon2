package isucon2

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func InitDb() {
	log.Println("Initializing database")
	f, err := os.Open(Conf.dir + "database/initial_data.sql")
	if err != nil {
		log.Panic(err.Error())
	}
	s := bufio.NewScanner(f)
	for s.Scan() {
		Db.Exec(s.Text())
	}

	if err := s.Err(); err != nil {
		log.Panic(err.Error())
	}
}
