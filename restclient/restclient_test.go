package restclient

import (
	"testing"
)

func TestGet(t *testing.T) {
	var s struct {
		UserID int `json:"userId"`
	}

	err := Get("http://jsonplaceholder.typicode.com/posts/1", &s)
	if err != nil {
		t.Errorf("Unexepected error %v", err)
	}

	if s.UserID != 1 {
		t.Errorf("Expected ID to be 1")
	}
}
