package fins

import (
	"os"
	"path/filepath"

	"github.com/jarosser06/fins/supermarket"
	chef "github.com/marpaia/chef-golang"
)

// TODO: Implement caching
var CacheFile string = filepath.Join(os.Getenv("HOME"), ".fins_cache.json")

type (
	ServerCookbooks      map[string]string
	ServerEnvironments   map[string]chef.Environment
	SupermarketCookbooks map[string]supermarket.Cookbook
)

type RemoteCache struct {
	superMarketEndpoint        string               `json:"supermarket_endpoint"`
	chefServerEndpoint         string               `json:"chef_server_endpoint"`
	timestamp                  string               `json:"timestamp"`
	latestServerCookbooks      ServerCookbooks      `json:"server_cookbooks"`
	latestSupermarketCookbooks SupermarketCookbooks `json:"supermarket_cookbooks"`
	environments               ServerEnvironments   `json:"environments"`
}

// Loads a fresh(Less than 10 minutes old) cache from disk if one exists
// does not error if one is not found
func (f *Fins) LoadCache() {
	log.Debug("LoadCache() not implemented")
}

func (f *Fins) ChefServerCookbooks() ServerCookbooks {
	cookbooks, err := f.chefClient.GetCookbooks()
	if err != nil {
		log.Fatal("error fetching cookbooks: %v", err)
	}

	cVers := make(map[string]string)
	for name, cookbook := range cookbooks {
		if len(cookbook.Versions) > 0 {
			cVers[name] = cookbook.Versions[0].Version
		} else {
			log.Warning("cookbook %s returned with no vesions")
		}
	}

	return cVers
}

// Returns a map of all latest cookbooks from supermarket that match
// the name on the chef server
func (f *Fins) LatestSupermarketCookbooks() SupermarketCookbooks {

	cookbooks := make(SupermarketCookbooks)
	for _, name := range f.ChefServerCookbooks() {
		c, err := f.supermarketClient.CookbookVersion(name, "latest")
		if err != nil {
			log.Error("%v", err)
		}
		cookbooks[name] = *c
		log.Debug("Supermarket Cookbook %s returned", name)
	}
	return cookbooks
}

func (f *Fins) LatestSupermarketCookbook(name string) supermarket.Cookbook {
	c, err := f.supermarketClient.CookbookVersion(name, "latest")
	if err != nil {
		log.Error("%v", err)
	}

	return *c
}

func (f *Fins) Environment(name string) (chef.Environment, bool) {
	e, ok, err := f.chefClient.GetEnvironment(name)
	if err != nil {
		log.Error("%v", err)
	}

	if !ok {
		log.Debug("environment %s not found", name)
		return chef.Environment{}, ok
	}

	return *e, ok
}
