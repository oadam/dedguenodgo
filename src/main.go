package nonoel

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()
var authRouter = router.PathPrefix("/authenticated-resources").Subrouter()
var unAuthRouter = router.PathPrefix("/unauthenticated-resources").Subrouter()

func init() {
	http.Handle("/", router)
	//auth
	authRouter.HandleFunc("/toto", wakeUp)
	//unauth
	unAuthRouter.HandleFunc("/wakeUp", wakeUp)
	unAuthRouter.HandleFunc("/party", createParty).Methods("POST")
}

func checkAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var partyId = request.Header.Get("dedguenodgo-partyId")
		//var partyPassword = request.Header.Get("dedguenodgo-partyPassword")
		h(w, r)
	}
}

func wakeUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

func createParty(w http.ResponseWriter, r *http.Request) {
	adminPassword := r.FormValue("adminPassword")
	partyName := r.FormValue("partyName")
	partyPassword := r.FormValue("partyPassword")

	fmt.Fprint(w, adminPassword)
	fmt.Fprint(w, partyName)
	fmt.Fprint(w, partyPassword)
}
