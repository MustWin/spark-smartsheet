package infrastructure

import (
	"core/domain"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

const userEmail = "foo@bar.baz"
const userStoreStr = `{
  "foo@bar.baz": {
    "Email": "foo@bar.baz",
    "Tokens": {
      "api": %q
    }
  }
}`

var users = domain.Users{}

func TestUserStore(t *testing.T) {
	dir, err := ioutil.TempDir("", "testuserstore")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	fp := path.Join(dir, "userstore.json")

	s, err := NewFileUserStore(fp)
	if err != nil {
		t.Error(err)
	}

	token, err := s.Users().RegisterUser(userEmail)
	if err != nil {
		t.Error(err)
	}

	t.Log(token)

	if err = s.Save(); err != nil {
		t.Error(err)
	}

	content, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Error(err)
	}

	expected, got := fmt.Sprintf(userStoreStr, token), string(content)
	if expected != got {
		t.Errorf("expected:\n%s\ngot:\n%s\n", expected, got)
	}
}
