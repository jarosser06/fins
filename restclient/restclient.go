package restclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Provides a simple method to handle both an
// http request and the JSON unmarshalling.
func Get(uri string, v interface{}) error {
	res, err := http.Get(uri)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	contents, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(contents, v)
	if err != nil {
		return err
	}

	return nil
}
