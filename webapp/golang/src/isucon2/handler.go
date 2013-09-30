package isucon2

import (
	"html/template"
	"net/http"
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
		log.Panic("Invalid artist_id:", artist_id)
	}

	data := map[string]interface{}{
		"Artists": GetArtist(artistId),
	}
	artistTmpl.ExecuteTemplate(w, "layout", data)
}
