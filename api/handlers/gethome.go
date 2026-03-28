package handlers

import (
	"itruj/discens-vestigo/api/middleware"
	"itruj/discens-vestigo/api/store"
	"itruj/discens-vestigo/api/templates"
	"net/http"
)

type HomeHandler struct{}

func NewHomeHandler() *HomeHandler {
	return &HomeHandler{}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user, ok := r.Context().Value(middleware.UserKey).(*store.User)

	if !ok {
		c := templates.GuestIndex()

		err := templates.Layout(c, "Item Aggregator").Render(r.Context(), w)

		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}

		return
	}

	c := templates.Index(user.Email)
	err := templates.Layout(c, "Getting Things Done").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
