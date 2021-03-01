package artists

import (
	"encoding/json"
	"net/http"

	"../search"
)

// Handler handles all requests to /artists/
func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			artistsHandler(w, r)
		}
	})
}

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	search.FloydWarshall()
	json.NewEncoder(w).Encode("hello")
}
