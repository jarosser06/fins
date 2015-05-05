package supermarket

import (
	"fmt"

	"github.com/jarosser06/fins/restclient"
)

const ChefEndpoint = "https://supermarket.chef.io/api/v1"

type UserCookbooks struct {
	Owns         map[string]string `json:"owns"`
	Collaborates map[string]string `json:"collaborates"`
	Follows      map[string]string `json:"follows"`
}

type User struct {
	Username  string        `json:"username"`
	Name      string        `json:"name"`
	Company   string        `json:"company"`
	Github    []string      `json:"github"`
	Twitter   string        `json:"twitter"`
	IRC       string        `json:"irc"`
	Jira      string        `json:"jira"`
	Cookbooks UserCookbooks `json:"cookbooks"`
}

type Cookbook struct {
	Name            string            `json:"name"`
	Maintainer      string            `json:"maintainer"`
	Description     string            `json:"description"`
	Dependencies    map[string]string `json:"dependencies"`
	Category        string            `json:"category"`
	Cookbook        string            `json:"cookbook"`
	LatestVersion   string            `json:"latest_version"`
	License         string            `json:"license"`
	ExternalURL     string            `json:"external_url"`
	AverageRating   string            `json:"average_rating"`
	CreatedAt       string            `json:"created_at"`
	UpdatedAt       string            `json:"updated_at"`
	Deprecated      bool              `json:"depcrecated"`
	Version         string            `json:"version"`
	Versions        []string          `json:"versions"`
	Platforms       map[string]string `json:"platforms"`
	File            string            `json:"file"`
	TarballFileSize int               `json:"tarball_file_size"`
	Metrics         struct {
		Downloads struct {
			Total    int `json:"total"`
			Versions map[string]int
		} `json:"downloads"`
		Followers int
	} `json:"metrics"`
}

type Client struct {
	Endpoint string
}

func NewClient() Client {
	return Client{Endpoint: ChefEndpoint}
}

func (c *Client) Cookbook(name string) (*Cookbook, error) {
	var ckbk *Cookbook
	uri := fmt.Sprintf("%s/cookbooks/%s", c.Endpoint, name)
	err := restclient.Get(uri, &ckbk)
	if err != nil {
		return ckbk, err
	}

	return ckbk, nil
}

func (c *Client) CookbookVersion(name string, version string) (*Cookbook, error) {
	var ckbk *Cookbook
	uri := fmt.Sprintf("%s/cookbooks/%s/versions/%s", c.Endpoint, name, version)
	err := restclient.Get(uri, &ckbk)
	if err != nil {
		return ckbk, err
	}

	return ckbk, nil
}

func (c *Client) User(name string) (*User, error) {
	var u *User
	uri := fmt.Sprintf("%s/users/%s", c.Endpoint, name)
	err := restclient.Get(uri, &u)
	if err != nil {
		return u, err
	}

	return u, nil
}
