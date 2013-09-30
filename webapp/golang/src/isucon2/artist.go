package isucon2

import (
	"log"
)

type Artist struct {
	Id   int
	Name string
}

func GetAllArtist() []Artist {
	rows, err := Db.Query("SELECT * FROM artist")
	if err != nil {
		log.Panic(err.Error())
	}

	artists := []Artist{}
	var id int
	var name string
	for rows.Next() {
		if err := rows.Scan(&id, &name); err != nil {
			log.Panic(err.Error())
		}
		artists = append(artists, Artist{id, name})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return artists
}
