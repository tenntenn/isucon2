package isucon2

import (
	"database/sql"
	"log"
	"time"
)

type Stock struct {
	Id          int
	VariationId int
	SeatId      string
	OrderId     int
	UpdatedAt   time.Time
}

func GetSeat(tx *sql.Tx, orderId, variationId int) *Stock {
	_, err := tx.Exec(`
        UPDATE
            stock
        SET
            order_id = ? 
        WHERE
            variation_id = ? AND
            order_id IS NULL
        ORDER BY
            RAND()
        LIMIT 1
    `, orderId, variationId)

	if err != nil {
		log.Panic(err.Error())
	}

	var rows *sql.Rows
	rows, err = tx.Query("SELECT id, seat_id, updated_at FROM stock WHERE order_id = ? LIMIT 1", orderId)
	if err != nil {
		log.Panic(err.Error())
	}

	if !rows.Next() {
		return nil
	}

	var id int
	var seatId string
	var updatedAt time.Time
	rows.Scan(&id, &seatId, &updatedAt)
	if err := rows.Err(); err != nil {
		log.Panic(err.Error())
	}

	return &Stock{id, variationId, seatId, orderId, updatedAt}
}
