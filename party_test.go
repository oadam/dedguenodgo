package nonoel

import (
	"appengine/aetest"
	"testing"
)

func TestParty(t *testing.T) {
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
	t.Logf("inexisting partyId")
	check := CheckPartyPassword(c, "nonexisting", "password")
	if check != false {
		t.Errorf("expecting party to not exist")
	}
	t.Logf("wrong password")
	check = CheckPartyPassword(c, "test", "wrongPassword")
	if check != false {
		t.Errorf("expecting password to be wrong")
	}
	t.Logf("good password")
	check = CheckPartyPassword(c, "test", "testPass")
	if check != true {
		t.Errorf("expecting password to be good")
	}
}
