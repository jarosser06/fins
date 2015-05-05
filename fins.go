package fins

import (
	"fmt"
	"os"
	"sync"

	"github.com/jarosser06/fins/pkg/version"
	"github.com/jarosser06/fins/supermarket"
	chef "github.com/marpaia/chef-golang"
	"github.com/op/go-logging"
)

var (
	logFormat = logging.MustStringFormatter("%{color:yellow} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}")
	log       = logging.MustGetLogger("fins")
)

// Main data structure for fins, keeps the config, client(s)
type Fins struct {
	cache             RemoteCache
	config            Config
	chefClient        chef.Chef
	supermarketClient supermarket.Client
}

// Returns a Fins struct
func Init(config string, logLevel string) Fins {
	c, err := LoadConfig(config)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	f := Fins{config: c}

	// Setup logging
	l, err := logging.LogLevel(logLevel)
	if err != nil {
		fmt.Println("Error: Bad log level ", logLevel)
	}
	b := logging.NewLogBackend(os.Stdout, "", 0)
	bFormatted := logging.NewBackendFormatter(b, logFormat)

	bLeveled := logging.AddModuleLevel(bFormatted)
	bLeveled.SetLevel(l, "")
	logging.SetBackend(bLeveled)

	// Get the chef client
	f.chefClient, err = c.ChefClient()
	if err != nil {
		log.Fatal("error creating ChefClient %v", err)
	}

	// Get the supermarket client
	f.supermarketClient = supermarket.NewClient()

	return f
}

// Displays Outdated Versions of Cookbooks in a Chef server
func (f *Fins) Outdated() int {
	var wg sync.WaitGroup
	var supCookbook supermarket.Cookbook

	for name, ver := range f.ChefServerCookbooks() {

		wg.Add(1)

		go func() {
			supCookbook = f.LatestSupermarketCookbook(name)
			wg.Done()
		}()
		wg.Wait()

		var supVer string
		if len(supCookbook.Version) == 0 {
			log.Debug("%s not found in supermarket", name)
			supVer = "not found"
			// Skip version compare
			goto L1
		} else {
			supVer = supCookbook.Version
		}

		if version.Compare(ver, supVer) != version.LessThan {
			continue
		}
	L1:
		fmt.Printf(
			"%s:\n  Chef Server: %s\n  Supermarket: %s\n",
			name,
			ver,
			supVer,
		)
	}

	return 0
}

// Displays a diff between the chef server and an environment
func (f *Fins) DiffServer(envName string) int {
	env, eOk := f.Environment(envName)
	if !eOk {
		fmt.Printf("Error: environment %s not found\n", envName)
		return 1
	}

	c := f.ChefServerCookbooks()
	if len(env.CookbookVersions) > 0 {
		for name, v := range env.CookbookVersions {

			// Handle when a cookbook isn't found
			var cVer string
			if ver, ok := c[name]; ok {
				cVer = ver
			} else {
				cVer = "no cookbook exists"
			}

			if !version.MatchConstraint(cVer, v) {
				fmt.Printf(
					"%s:\n  Environment Constraint: %s\n  Latest Cookbook Version: %s\n",
					name,
					v,
					cVer,
				)
			}
		}
	}

	return 0
}

// Displays a diff between two chef environments
func (f *Fins) DiffEnvironments(env1 string, env2 string) int {
	e1, eOk := f.Environment(env1)
	if !eOk {
		fmt.Printf("Error: environment %s not found\n", env1)
		return 1
	}

	e2, eOk := f.Environment(env2)
	if !eOk {
		fmt.Printf("Error: environment %s not found\n", env2)
		return 1
	}

	c := f.ChefServerCookbooks()
	for _, cName := range mergeCookbookLists(e1.CookbookVersions, e2.CookbookVersions) {
		//for cName, e1Ver := range e1.CookbookVersions {
		var e1Ver, e2Ver string
		if ver, ok := e1.CookbookVersions[cName]; ok {
			log.Debug("found cookbook constraint %s for %s", ver, cName)
			e1Ver = ver
		} else {
			log.Debug("%s has no constraint for %s", env2, cName)
			e1Ver = "no version constraint"
		}

		if ver, ok := e2.CookbookVersions[cName]; ok {
			log.Debug("found cookbook constraint %s for %s", ver, cName)
			e2Ver = ver
		} else {
			log.Debug("%s has no constraint for %s", env2, cName)
			e2Ver = "no version constraint"
		}

		if e1Ver == e2Ver {
			log.Debug("env \"%s\"@%s and env \"%s\"@%s constrains match", e1Ver, env1, e2Ver, env2)
			continue
		}

		var cVer string
		if ver, ok := c[cName]; ok {
			cVer = ver
		} else {
			cVer = "no cookbook exists"
		}

		fmt.Printf(
			"%s:\n  %s Constraint: %s\n  %s Constraint: %s\n  Latest Cookbook Version: %s\n",
			cName,
			env1,
			e1Ver,
			env2,
			e2Ver,
			cVer,
		)
	}

	return 0
}

// Take the Cookbook Version map and produce a slice of cookbook names
// to be iterated over
func mergeCookbookLists(list1 map[string]string, list2 map[string]string) []string {
	var result []string
	var wg sync.WaitGroup
	in := make(chan string)
	out := make(chan string)

	// Send cookbooks from list1 onto the channel
	wg.Add(2)
	go func() {
		for name, _ := range list1 {
			in <- name
		}
		wg.Done()
	}()

	go func() {
		for name, _ := range list2 {
			in <- name
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(in)
	}()

	// Check for duplicates and discard
	go func() {
		listMap := make(map[string]bool)
		for {
			if name, ok := <-in; ok {
				if _, k := listMap[name]; !k {
					listMap[name] = true
					out <- name
				}
			} else {
				break
			}
		}
		close(out)
	}()

	for {
		if name, ok := <-out; ok {
			result = append(result, name)
		} else {
			break
		}
	}

	return result
}
