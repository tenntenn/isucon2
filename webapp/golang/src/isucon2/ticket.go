package isucon2

import (
	"log"
)

type Ticket struct {
	Id       int
	Name     string
	ArtistId int
	Count    int
}

func ticketCount(ticketId int) int {
	row := Db.QueryRow(
		`SELECT COUNT(*) FROM variation
            INNER JOIN stock ON stock.variation_id = variation.id
            WHERE variation.ticket_id = ? AND stock.order_id IS NULL`,
		ticketId,
	)

	var count int
	if err := row.Scan(&count); err != nil {
		log.Panic(err.Error())
	}

	return count
}

func GetAllTickets(artistId int) []*Ticket {
	rows, err := Db.Query(
		"SELECT id, name FROM ticket WHERE artist_id = ?",
		artistId,
	)
	if err != nil {
		log.Panic(err.Error())
	}

	tickets := []*Ticket{}
	var id int
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Panic(err.Error())
		}
		tickets = append(tickets, &Ticket{id, name, artistId, ticketCount(id)})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return tickets
}
