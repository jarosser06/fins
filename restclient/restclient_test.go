package restclient

import (
	"testing"
)

func TestGet(t *testing.T) {
	var s struct {
		IP string `json:"ip"`
	}

	err := Get("http://ip.jsontest.com/", &s)
	if err != nil {
		t.Errorf("Unexepected error %v", err)
	}

	if s.IP == "" {
		t.Errorf("Expected IP to be populated")
	}
}
