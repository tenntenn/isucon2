package isucon2

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	indexTmpl    = parseTemplate("index")
	artistTmpl   = parseTemplate("artist")
	completeTmpl = parseTemplate("complete")
	soldoutTmpl  = parseTemplate("soldout")
)

func parseTemplate(name string) *template.Template {
	return template.Must(template.New(name).ParseFiles("templates/layout.html", "templates/"+name+".html"))

}

func TopPageHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Artists":    GetAllArtist(),
		"RecentSold": GetRecentSold(),
	}
	indexTmpl.ExecuteTemplate(w, "layout", data)
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	artistId, err := strconv.ParseInt(r.FormValue("artist_id"), 10, 64)
	if err != nil {
		log.Panicf("Invalid artist_id: %d", artistId)
	}

	data := map[string]interface{}{
		"Artists": GetArtist(int(artistId)),
	}
	artistTmpl.ExecuteTemplate(w, "layout", data)
}

func BuyHandler(w http.ResponseWriter, r *http.Request) {
	variationId, err := strconv.ParseInt(r.FormValue("variation_id"), 10, 64)
	if err != nil {
		log.Panic(err.Error())
	}
	memberId := r.FormValue("memmber_id")

	var tx *sql.Tx
	tx, err = Db.Begin()
	if err != nil {
		log.Panic(err.Error())
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	order := NewOrder(tx, memberId)
	stock := GetSeat(tx, order.Id, int(variationId))
	if stock == nil {
		tx.Rollback()
		soldoutTmpl.ExecuteTemplate(w, "layout", nil)
		return
	}

	tx.Commit()
	completeTmpl.ExecuteTemplate(w, "layout", stock)
}
