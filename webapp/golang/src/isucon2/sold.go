package isucon2

import (
	"log"
)

type Sold struct {
	SeatId        int
	VariationName string
	TicketName    string
	AritistName   string
}

func GetRecentSold() []Sold {
	rows, err := Db.Query(`
SELECT stock.seat_id, variation.name AS v_name, ticket.name AS t_name, artist.name AS a_name FROM stock
        JOIN variation ON stock.variation_id = variation.id
        JOIN ticket ON variation.ticket_id = ticket.id
        JOIN artist ON ticket.artist_id = artist.id
        WHERE order_id IS NOT NULL
        ORDER BY order_id DESC LIMIT 10
  `)
	if err != nil {
		log.Panic(err.Error())
	}

	var seatId int
	var vName, tName, aName string
	solds := []Sold{}
	for rows.Next() {
		if err := rows.Scan(&seatId, &vName, &tName, &aName); err != nil {
			log.Panic(err.Error())
		}
		solds = append(solds, Sold{seatId, vName, tName, aName})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return solds
}
