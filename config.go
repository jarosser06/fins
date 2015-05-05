package fins

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/jarosser06/fins/pkg/fileutil"
	chef "github.com/marpaia/chef-golang"
)

type Config struct {
	Chef struct {
		ChefServerUrl string `json:"chef_server_url"`
		NodeName      string `json:"node_name"`
		ClientKey     string `json:"client_key"`
		Version       string `json:"version"`
	} `json:"chef"`
	Supermarket struct {
		Endpoint string `json:"endpoint"`
	} `json:"supermarket"`
	SSL struct {
		Verify bool `json:"verify"`
	} `json:"ssl"`
}

// Loads the fins json config file and returns
// a populated Config struct.
func LoadConfig(config string) (Config, error) {
	var c Config

	// Return an error if the file doesn't exist
	if !fileutil.FileExist(config) {
		return c, fmt.Errorf("file %s does not exist", config)
	}

	f, err := ioutil.ReadFile(config)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(f, &c)
	if err != nil {
		return c, err
	}

	// Allows the relative path for a client key to be passed
	if !filepath.IsAbs(c.Chef.ClientKey) {
		confAbs, _ := filepath.Abs(config)
		c.Chef.ClientKey = filepath.Join(filepath.Dir(confAbs), c.Chef.ClientKey)
	}

	return c, c.validate()
}

// Private method used when calling LoadConfig to
// validate the minimum information is provided in
// the config.
func (c *Config) validate() error {
	switch {
	case c.Chef.ChefServerUrl == "":
		return errors.New("chef_server_url cannot be empty")
	case c.Chef.NodeName == "":
		return errors.New("node_name cannot be empty")
	case c.Chef.ClientKey == "":
		return errors.New("client_key cannot be empty")
	}
	return nil
}

// Returns a Chef struct based on information in the
// Config struct.
func (c *Config) ChefClient() (chef.Chef, error) {
	chefAPI := chef.Chef{
		Url:         c.Chef.ChefServerUrl,
		UserId:      c.Chef.NodeName,
		SSLNoVerify: !c.SSL.Verify,
	}

	if c.Chef.Version == "" {
		chefAPI.Version = "11.6.0"
	} else {
		chefAPI.Version = c.Chef.Version
	}

	key, err := keyFromFile(c.Chef.ClientKey)
	if err != nil {
		return chefAPI, err
	}

	chefAPI.Key = key

	return chefAPI, nil
}

func keyFromFile(filename string) (*rsa.PrivateKey, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return keyFromString(content)
}

func keyFromString(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, fmt.Errorf("block size invalid for '%s'", string(key))
	}
	rsaKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsaKey, nil
}
