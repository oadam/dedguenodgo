package nonoel

import (
	"appengine"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	r := mux.NewRouter()
	http.Handle("/", r)
	main := r.PathPrefix("/REST/").Subrouter()
	RegisterUserGroupHandlers(main)
}

func AddAuth(subpath string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var partyId = r.Header.Get("dedguenodgo-partyId")
		var partyPassword = r.Header.Get("dedguenodgo-partyPassword")
		if !checkAuth(partyId, partyPassword) {
			http.Error(w, "wrong username/password", http.StatusUnauthorized)
		} else {
			handler.ServeHTTP(w, r)
		}
	})
}

func checkAuth(partyId, partyPassword string) bool {
	return partyId == "partyId" && partyPassword == "partyPassword"
}

func noop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL)
	fmt.Fprint(w, " path ")
	fmt.Fprint(w, r.URL.Path)
}
