package nonoel

import (
	"appengine"
	"appengine/datastore"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalid = bcrypt.ErrMismatchedHashAndPassword

type adminPassword struct {
	password []byte
}

func CheckOrSetAdminPassword(c appengine.Context, password string) error {
	key := datastore.NewKey(c, "adminPassword", "singleton", 0, nil)
	var current *adminPassword = new(adminPassword)
	err := datastore.Get(c, key, current)
	switch err {
	case nil:
		// check password
		fmt.Printf("current: %v", current.password)
		return bcrypt.CompareHashAndPassword(current.password, []byte(password))
	case datastore.ErrNoSuchEntity:
		// no current password
		// store new password and exit
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		_, err = datastore.Put(c, key, &adminPassword{hashed})
		return err
	default:
		return err
	}
}
