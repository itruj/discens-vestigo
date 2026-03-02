package api

import (
	// "html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func (s *Server) handleDefineWorkView(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.ParseFiles("./static/define-work.html")
	// err = tmpl.Execute(w, data)

	// if err != nil {
	// 	http.Error(w, "Error executing template", http.StatusInternalServerError)
	//	return
	// }
}
