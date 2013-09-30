package isucon2

import (
	"html/template"
	"net/http"
)

var (
	indexTmpl = parseTemplate("index")
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
