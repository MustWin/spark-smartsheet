package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"core/domain"
)

type UserStore interface {
	Load() error
	Save() error
	Users() domain.Users
}

type FileUserStore struct {
	once  sync.Once
	path  string
	users domain.Users
}

// NewFileUserStore creates a UserStore from the given path
func NewFileUserStore(path string) (*FileUserStore, error) {
	s := &FileUserStore{path: path, users: domain.Users{}}
	return s, s.Load()

}

// Save the UserStore to the configured filepath
func (s *FileUserStore) Save() error {
	b, err := json.MarshalIndent(s.users, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(s.path, b, os.ModePerm)
}

// Load populates the UserStore from the configured filepath
func (s *FileUserStore) Load() error {
	s.once.Do(func() { initializeIfNotExists(s.path, "{}") })
	b, err := ioutil.ReadFile(s.path)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, &s.users)
}

// Users return current users
func (s *FileUserStore) Users() domain.Users { return s.users }

func initializeIfNotExists(path string, contents string) error {
	var err error
	if _, err = os.Stat(path); os.IsNotExist(err) {
		var handle *os.File
		if handle, err = os.Create(path); err == nil {
			if _, err = handle.Write([]byte(contents)); err == nil {
				return handle.Close()
			}
		}
	}
	return err
}
