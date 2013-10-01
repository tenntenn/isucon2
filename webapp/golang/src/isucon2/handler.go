package isucon2

import (
	"database/sql"
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var (
	indexTmpl    = parseTemplate("index")
	artistTmpl   = parseTemplate("artist")
	completeTmpl = parseTemplate("complete")
	soldoutTmpl  = parseTemplate("soldout")
	adminTmpl    = parseTemplate("admin")
	ticketTmpl   = parseTemplate("ticket")
)

var idRegexp = regexp.MustCompile(".+/([0-9]+)")

func getId(path string) (int, error) {
	id, err := strconv.ParseInt(idRegexp.FindStringSubmatch(path)[1], 10, 64)
	return int(id), err
}

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

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artistId, err := getId(r.RequestURI)
	if err != nil {
		log.Panicf("Invalid artistId: %d", artistId)
	}

	data := map[string]interface{}{
		"Artists":    GetArtist(int(artistId)),
		"Tickets":    GetAllTickets(int(artistId)),
		"RecentSold": GetRecentSold(),
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

	data := map[string]interface{}{
		"Stock":      stock,
		"RecentSold": GetRecentSold(),
	}
	completeTmpl.ExecuteTemplate(w, "layout", data)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		InitDb()
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	adminTmpl.ExecuteTemplate(w, "layout", nil)
}

func AdminCsvHandler(w http.ResponseWriter, r *http.Request) {
	wr := csv.NewWriter(w)
	for _, order := range GetAllOrder() {
		err := wr.Write([]string{
			strconv.Itoa(order.Id),
			order.MemberId,
			order.Stock.SeatId,
			strconv.Itoa(order.Stock.VariationId),
			order.Stock.UpdatedAt.Format("2013-10-01 13:11:11"),
		})
		if err != nil {
			log.Panic(err.Error())
		}
	}

	wr.Flush()
	w.Header().Set("Content-type", "text/csv")
}

func TicketHandler(w http.ResponseWriter, r *http.Request) {
	ticketId := getId(r.RequestURI)
	ticket := GetTicket(ticketId)
	variations := GetVariations()
}
