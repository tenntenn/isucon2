package isucon2

import (
	"database/sql"
	"log"
)

type Order struct {
	Id       int
	MemberId string
}

func NewOrder(tx *sql.Tx, memberId string) *Order {
	r, err := tx.Exec("INSERT INTO order_request (member_id) VALUES (?)", memberId)
	if err != nil {
		log.Panic(err.Error())
	}

	id, err := r.LastInsertId()
	if err != nil {
		log.Panic(err.Error())
	}

	return &Order{int(id), memberId}
}
