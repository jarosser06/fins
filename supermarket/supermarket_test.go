package supermarket

import "testing"

func TestCookbook(t *testing.T) {
	c := NewClient()

	cName := "sysdig"

	ckbk, err := c.Cookbook(cName)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if ckbk.Name != cName {
		t.Errorf("Expected cookbook name %s not %s", cName, ckbk.Name)
	}
}

func TestCookbookVersion(t *testing.T) {
	c := NewClient()
	cName := "sysdig"
	cVer := "0.3.1"

	ckbk, err := c.CookbookVersion(cName, cVer)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if ckbk.Version != cVer {
		t.Errorf("Expected cookbook version %s not %s", cVer, ckbk.Version)
	}
}

func TestUser(t *testing.T) {
	c := NewClient()
	uName := "jarosser06"

	u, err := c.User(uName)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}

	if u.Name != "Jim Rosser" {
		t.Errorf("Expected name Jim Rosser not %s", u.Name)
	}
}
