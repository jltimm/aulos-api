package artists

import "net/http"

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
	w.Write([]byte("You've hit the artists endpoint!"))
}
