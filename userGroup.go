package nonoel

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterUserGroupHandlers(h *mux.Router) {
	userGroup := h.PathPrefix("/userGroup/").Subrouter()
	userGroup.Methods('PUT').Handler("/", httfunc(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.URL.Path)
	})
	h.HandleFunc("/userGroup/toto", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "yomama")
	})
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
	fmt.Fprintf(w, "Party %s/%s created successfully", partyName, partyPassword)
}
