package isucon2

import (
	"log"
)

type Artist struct {
	Id   int
	Name string
}

func GetArtist(artistId int) *Artist {
	row := Db.QueryRow(
		"SELECT id, name FROM artist WHERE id = ? LIMIT 1",
		artistId,
	)

	var id int
	var name string
	if err := row.Scan(&id, &name); err != nil {
		log.Panic(err.Error())
	}

	return &Artist{id, name}
}

func GetAllArtist() []*Artist {
	rows, err := Db.Query("SELECT * FROM artist")
	if err != nil {
		log.Panic(err.Error())
	}

	artists := []*Artist{}
	var id int
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Panic(err.Error())
		}
		artists = append(artists, &Artist{id, name})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return artists
}
