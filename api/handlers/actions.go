package api

import (
	"fmt"
	"html/template"
	"itruj/discens-vestigo/api/store"
	"log"
	"net/http"
)

func (s *Server) handleActionsView(w http.ResponseWriter, r *http.Request) {
	rows, err := s.DB.Model(&store.Card{}).Select("id", "title", "done").Rows()

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var list []store.Card
	for rows.Next() {
		var t store.Card
		rows.Scan(&t.ID, &t.Title, &t.Done)
		fmt.Printf("%+v\n", t)
		list = append(list, t)
	}

	data := struct {
		Cards []store.Card
	}{
		Cards: list,
	}

	tmpl, err := template.ParseFiles("./static/actions.html")
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
