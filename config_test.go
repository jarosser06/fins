package fins

import "testing"

func TestLoadConfig(t *testing.T) {
	c, err := LoadConfig("tests/fins.json")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	testURL := "https://localhost:8443/organizations/testorg"
	if c.Chef.ChefServerUrl != testURL {
		t.Errorf("Expectetd ChefServerURL to be %s", testURL)
	}
}

func TestChefClient(t *testing.T) {
	c, err := LoadConfig("tests/fins.json")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	client, _ := c.ChefClient()
	_, err = client.GetCookbooks()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
