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
	RegisterPartyHandlers(main)
}

func AddAuth(handler AuthenticatedHandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := appengine.NewContext(r)
		var partyId = r.Header.Get("dedguenodgo-partyId")
		var partyPassword = r.Header.Get("dedguenodgo-partyPassword")
		if partyId == "" || partyPassword == "" {
			http.Error(w, "received empty id or password", http.StatusUnauthorized)
			return
		}
		if !CheckPartyPassword(c, partyId, partyPassword) {
			http.Error(w, "wrong username/password", http.StatusUnauthorized)
		} else {
			handler(partyId, w, r)
		}
	})
}

type AuthenticatedHandlerFunc func(partyName string, w http.ResponseWriter, r *http.Request)

func noop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.URL)
	fmt.Fprint(w, " path ")
	fmt.Fprint(w, r.URL.Path)
}
