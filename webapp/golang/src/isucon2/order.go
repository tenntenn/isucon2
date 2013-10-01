package isucon2

import (
	"database/sql"
	"log"
)

type Order struct {
	Id       int
	MemberId string
	Stock    *Stock
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

func GetAllOrder() *[]Order {
	rows, err := Db.Query(`
  SELECT
    order_request.*, stock.id, stock.seat_id, stock.variation_id, stock.updated_at
  FROM
    order_request
  JOIN stock
  ON
    order_request.id = stock.order_id
  ORDER BY
    order_request.id ASC
  `)
	if err != nil {
		log.Panic(err.Error())
	}

	orders := []*Order{}

	var (
		oid      int
		memberId string
	)

	var (
		sid         int
		seatId      string
		variationId int
		updatedAt   time.Time
	)

	for rows.Next() {
		rows.Scan(&oid, &memberId, &sid, &seatId, &variationId, &updatedAt)
		stock := &Stock{sid, variationId, seatId, orderId, updatedAt}
		orders = append(orders, &Order{&oid, &memberId, stock})
	}

	if err := rows.Err(); err != nil {
		log.Panic(err.Error())
	}

	return orders
}
