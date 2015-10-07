package nonoel

import (
	"appengine/aetest"
	"testing"
)

type Toto struct {
	Name string
}

func TestAnonymousUsers(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	t.Logf("creating party")
	err = storeParty(c, "test", "testPass")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("fetch users")
	users, err := fetchPartyUsers(c, "test")
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 0 {
		t.Error("found more than 0 users in test party")
	}

	t.Logf("add user")
	user := User{Name: "testuser", Id: ""}
	err = putAnonymousPartyUser(c, "test", &user)
	if err != nil {
		t.Fatal(err)
	}
	if user.Id == 0 {
		t.Fatal("user Id not specified after saving")
	}

	t.Logf("fetch added user")
	users, err = fetchPartyUsers(c, "test")
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 {
		t.Errorf("expected 1 users, received :%v", users)
	}
}
