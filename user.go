package nonoel

import (
	"appengine"
	"appengine/datastore"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

const membershipKind = "party-membership"

func RegisterUserHandlers(h *mux.Router) {
	h.Path("/party-users").Methods("POST").Handler(AddAuth(getPartyUsers))
	h.Path("/user").Methods("POST").Handler(AddAuth(addAnonymousPartyUser))
}

type User struct {
	Id   string `json:"id" datastore:"-"`
	Name string `json:"name"`
}

// Party key is used as the ancestor key
type PartyMembership struct {
	User *datastore.Key
}

func getUserKey(c appengine.Context, id int64) *datastore.Key {
	return datastore.NewKey(c, "user", "", id, nil)
}

func getPartyUsers(partyId string, w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	users, err := fetchPartyUsers(c, partyId)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		c.Errorf("error while fetching party members : %v", err)
		return
	}
	result, err := json.Marshal(users)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		c.Errorf("error while converting party members to json : %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func addAnonymousPartyUser(partyId string, w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		http.Error(w, "error", http.StatusBadRequest)
		c.Warningf("error while parsing user : %v", err)
		return
	}
	if user.Name == "" {
		http.Error(w, "user name cannot be empty", http.StatusBadRequest)
		return
	}
	if user.Id != "" {
		http.Error(w, "user id should be unspecified", http.StatusBadRequest)
		c.Warningf("user id should be unspecified : %v", user)
		return
	}
	err = putAnonymousPartyUser(c, partyId, &user)
	result, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "error", http.StatusInternalServerError)
		c.Errorf("error while converting user to json : %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func fetchPartyUsers(c appengine.Context, partyId string) ([]User, error) {
	partyKey := GetPartyKey(c, partyId)
	q := datastore.NewQuery(membershipKind).
		Ancestor(partyKey)
	var memberships []PartyMembership
	_, err := q.GetAll(c, &memberships)
	if err != nil {
		return nil, err
	}
	var userKeys []*datastore.Key
	for _, membership := range memberships {
		userKeys = append(userKeys, membership.User)
	}
	result := make([]User, len(userKeys))
	err = datastore.GetMulti(c, userKeys, result)
	return result, err
}

func putAnonymousPartyUser(c appengine.Context, partyId string, user *User) error {
	k := datastore.NewIncompleteKey(c, "user", nil)
	userKey, err := datastore.Put(c, k, user)
	if err != nil {
		return err
	}
	user.Id = strconv.FormatInt(userKey.IntID(), 10)
	partyKey := GetPartyKey(c, partyId)
	var membership = new(PartyMembership)
	membership.User = userKey
	mk := datastore.NewIncompleteKey(c, membershipKind, partyKey)
	_, err = datastore.Put(c, mk, membership)
	return err
}
