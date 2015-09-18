package nonoel

import (
	"appengine"
	"fmt"
	"net/http"
)

func init() {
	addUnauth("/wakeUp", noop)
	addUnauth("/party", createParty)
	addAuth("/hello", noop)
}

func addUnauth(subpath string, f func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc("/unauthenticated-resources"+subpath, f)
}

func addAuth(subpath string, f func(w http.ResponseWriter, r *http.Request)) {
	http.Handle("/authenticated-resources"+subpath, addAuthCheck(http.HandlerFunc(f)))
}

func addAuthCheck(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var partyId = r.Header.Get("dedguenodgo-partyId")
		var partyPassword = r.Header.Get("dedguenodgo-partyPassword")
		if !checkAuth(partyId, partyPassword) {
			http.Error(w, "wrong username/password", http.StatusUnauthorized)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

func checkAuth(partyId, partyPassword string) bool {
	return partyId == "partyId" && partyPassword == "partyPassword"
}

func noop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func createParty(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	adminPassword := r.FormValue("adminPassword")
	err := CheckOrSetAdminPassword(c, adminPassword)
	if err != nil {
		http.Error(w, "wrong admin password", http.StatusUnauthorized)
		return
	}

	partyName := r.FormValue("partyName")
	partyPassword := r.FormValue("partyPassword")

	fmt.Fprintf(w, "Party %s/%s created successfully", partyName, partyPassword)
}
