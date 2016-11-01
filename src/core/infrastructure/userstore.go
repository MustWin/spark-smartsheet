package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"core/domain"
)

// UserStore marshals the Users data structure to and from a JSON
// formatted file
type UserStore struct {
	domain.Users
	path string
}

// NewFileUserStore creates a UserStore from the given path
func NewFileUserStore(path string) (*UserStore, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if h, e2 := os.Create(path); e2 != nil {
			return nil, e2
		} else {
			if _, e2 = h.Write([]byte(`{}`)); e2 != nil {
				return nil, e2
			}
			if e2 = h.Close(); e2 != nil {
				return nil, e2
			}
		}
	} else if err != nil {
		return nil, err
	}
	s := &UserStore{path: path}
	return s, s.Load()

}

// Load populates the UserStore from the configured filepath
func (s *UserStore) Save() error {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(s.path, b, os.ModePerm)
}

// Save the UserStore to the configured filepath
func (s *UserStore) Load() error {
	b, err := ioutil.ReadFile(s.path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, s)
}
