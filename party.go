package nonoel

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func RegisterPartyHandlers(h *mux.Router) {
	h.Path("/party").Methods("POST").Handler(http.HandlerFunc(createParty))
	//h.Path("/party-users").Methods("POST").Handler(http.HandlerFunc(getPartyUsers))
}

func createParty(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	adminPassword := r.FormValue("adminPassword")
	err := CheckOrSetAdminPassword(c, adminPassword)
	if err != nil {
		switch err {
		case ErrInvalidPassword:
			http.Error(w, "wrong admin password", http.StatusUnauthorized)
			c.Warningf("wrong admin password")
		default:
			http.Error(w, "error", http.StatusInternalServerError)
			c.Errorf("error while checking adminPassword : %v", err)
		}
		return
	}
	partyName := r.FormValue("partyName")
	partyPassword := r.FormValue("partyPassword")
	err = storeParty(c, partyName, partyPassword)
	if err != nil {
		http.Error(w, "wrong username/password", http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "Party %s created successfully", partyName)
	}
}

type Party struct {
	Password []byte
}

func storeParty(c appengine.Context, name, password string) error {
	key := datastore.NewKey(c, "party", name, 0, nil)
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	var toStore = &Party{hashed}
	_, err = datastore.Put(c, key, toStore)
	return err
}
