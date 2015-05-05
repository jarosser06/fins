package fins

import (
	"testing"
)

func TestMergeCookbookList(t *testing.T) {
	res := mergeCookbookLists(
		map[string]string{"yum": "", "chef-sugar-rackspace": "", "mysql": ""},
		map[string]string{"yum": "", "postgresql": "", "database": ""},
	)

	if len(res) != 5 {
		t.Errorf("expected resulting list to be 5")
	}
}
