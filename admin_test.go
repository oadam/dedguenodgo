package nonoel

import (
	"appengine/aetest"
	"testing"
)

func TestAdminPassword(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	err = CheckOrSetAdminPassword(c, "testPassword")
	if err != nil {
		t.Fatal(err)
	}
	err = CheckOrSetAdminPassword(c, "testPassword")
	if err != nil {
		t.Fatalf("got an error while checking password %v", err)
	}
	err = CheckOrSetAdminPassword(c, "wrongPassword")
	if err == nil {
		t.Fatal("did not get an error while trying wrong password")
	}
}
