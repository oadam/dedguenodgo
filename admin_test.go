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

	t.Logf("setting password")
	err = CheckOrSetAdminPassword(c, "testPassword")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("checking password")
	err = CheckOrSetAdminPassword(c, "testPassword")
	if err != nil {
		t.Fatalf("got an error while checking password %v", err)
	}
	t.Logf("trying wrong password")
	err = CheckOrSetAdminPassword(c, "wrongPassword")
	if err == nil {
		t.Fatal("did not get an error while trying wrong password")
	}
}
