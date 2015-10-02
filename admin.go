package nonoel

import (
	"appengine"
	"appengine/datastore"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidPassword = bcrypt.ErrMismatchedHashAndPassword

type adminPassword struct {
	Password []byte
}

func CheckOrSetAdminPassword(c appengine.Context, password string) error {
	key := datastore.NewKey(c, "adminPassword", "singleton", 0, nil)
	current := new(adminPassword)
	var err = datastore.Get(c, key, current)
	switch err {
	case nil:
		// check password
		return bcrypt.CompareHashAndPassword(current.Password, []byte(password))
	case datastore.ErrNoSuchEntity:
		// no current password
		// store new password and exit
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		var toStore = &adminPassword{hashed}
		_, err = datastore.Put(c, key, toStore)
		return err
	default:
		return err
	}
}
