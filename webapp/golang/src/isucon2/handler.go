package isucon2

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var (
	indexTmpl  = parseTemplate("index")
	artistTmpl = parseTemplate("artist")
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
