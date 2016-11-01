package infrastructure

import (
	"core/domain"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
)

const userStoreStr = `{
  "Users": null
}`

var users = domain.Users{}

func TestUserStore(t *testing.T) {
	dir, err := ioutil.TempDir("", "testuserstore")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	fp := path.Join(dir, "userstore.json")

	s, err := NewFileUserStore(fp)
	if err != nil {
		t.Error(err)
	}

	if err = s.Save(); err != nil {
		t.Error(err)
	}

	content, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Error(err)
	}

	if string(content) != userStoreStr {
		t.Error(string(content))
	}
}
