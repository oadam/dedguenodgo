package nonoel

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	authPrefix   = "/authenticated-resources/"
	unAuthPrefix = "/unauthenticated-resources/"
)

var (
	unAuthRouter = mux.NewRouter()
	authRouter   = mux.NewRouter()
)

func init() {
	http.Handle(unAuthPrefix, http.StripPrefix(unAuthPrefix, unAuthRouter))
	http.Handle(authPrefix, addAuthCheck(http.StripPrefix(authPrefix, authRouter)))
	//auth
	authRouter.HandleFunc("/toto", wakeUp)
	authRouter.HandleFunc("/{yo:.*}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL)
	})
	//unauth
	unAuthRouter.HandleFunc("/wakeUp", wakeUp)
	unAuthRouter.HandleFunc("/", wakeUp)
	unAuthRouter.HandleFunc("/{yo:.*}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL)
	})
	//unAuthRouter.HandleFunc("/party", createParty).Methods("POST")
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
	return partyId != "partyId" || partyPassword != "partyPassword"
}

func wakeUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

/*func createParty(w http.ResponseWriter, r *http.Request) {
	adminPassword := r.FormValue("adminPassword")
	partyName := r.FormValue("partyName")
	partyPassword := r.FormValue("partyPassword")

	fmt.Fprintf(w, "Party %s created successfully", partyName)
}*/
