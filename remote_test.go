package fins

import (
	"testing"

	"github.com/op/go-logging"
)

const TestConfig = "./tests/fins.json"

func finsTestInit(cache bool) Fins {
	c, _ := LoadConfig(TestConfig)
	return Init(c, logging.ERROR, cache)
}

func TestChefServerCookbooks(t *testing.T) {
	f := finsTestInit(false)
	if c := f.ChefServerCookbooks(); !(len(c) > 0) {
		t.Errorf("expected cookbook list to be greater than 0")
	}
}

/*
func TestLatestSupermarketCookbooks(t *testing.T) {
	f := finsTestInit(false)
	c := f.LatestSupermarketCookbooks()
	if !(len(c) > 0) {
		t.Errorf("expected latest cookbooks list to be greather than 0")
	}

	if _, ok := c["yum"]; !ok {
		t.Errorf("expected yum cookbook to be in the returned list")
	}
}
*/

func TestLatestSupermarketCookbook(t *testing.T) {
	f := finsTestInit(false)
	c := f.LatestSupermarketCookbook("yum")
	if c.License != "Apache 2.0" {
		t.Errorf("expected yum cookbook to have Apache 2.0 license")
	}
}

func TestEnvironment(t *testing.T) {
	f := finsTestInit(false)
	e, _ := f.Environment("staging")
	if !(len(e.CookbookVersions) > 0) {
		t.Errorf("expected staging cookbook versions to be greater than 0")
	}
}
