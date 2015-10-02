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
	h.Path("/party-users").Methods("POST").Handler(AddAuth(getPartyUsers))
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

func getPartyUsers(partyId string, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "[]")
}

func CheckPartyPassword(c appengine.Context, id, password string) bool {
	key := getPartyKey(c, id)
	current := new(Party)
	var err = datastore.Get(c, key, current)
	passOk := errBoolOrPanic(err, nil, datastore.ErrNoSuchEntity)
	if !passOk {
		return false
	}
	err = bcrypt.CompareHashAndPassword(current.Password, []byte(password))
	return errBoolOrPanic(err, nil, bcrypt.ErrMismatchedHashAndPassword)
}

func errBoolOrPanic(err, trueErr, falseErr error) bool {
	switch err {
	case trueErr:
		return true
	case falseErr:
		return false
	default:
		panic(err)
	}
}

type Party struct {
	Password []byte
}

func getPartyKey(c appengine.Context, id string) *datastore.Key {
	return datastore.NewKey(c, "party", id, 0, nil)
}

func storeParty(c appengine.Context, id, password string) error {
	key := getPartyKey(c, id)
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	var toStore = &Party{hashed}
	_, err = datastore.Put(c, key, toStore)
	return err
}
