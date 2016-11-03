package domain

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"sort"
)

// Errors relating to User
var (
	ErrorUserExists    = errors.New("user exists")
	ErrorInvalidEmail  = errors.New("invalid email")
	ErrorInvalidToken  = errors.New("invalid token")
	ErrorInternalError = errors.New("internal error")
)

// User encapsulates user specific info
type User struct {
	Email  string
	Tokens map[string]string
}

// Users is the collection of all users
type Users map[string]User

// RegisterUser creates a new access token for a given email address
func (u Users) RegisterUser(email string) (string, error) {
	if ok := validate(email); !ok {
		return "", ErrorInvalidEmail
	}
	if _, ok := u[email]; ok {
		return "", ErrorUserExists
	}
	token, err := generateRandomString(email, 32)
	if err != nil {
		log.Printf("registering %q; error generating random string: %v", email, err)
		return "", ErrorInternalError
	}
	u[email] = User{Email: email, Tokens: map[string]string{"api": token}}
	return token, nil
}

// VerifyUser verifies that a user provides the correct key
func (u Users) VerifyUser(token string) bool {
	email, err := extractPrefix(token)
	if err != nil {
		log.Printf("error extracting prefix from token %q: %v", token, err)
		return false
	}
	if user, ok := u[email]; ok {
		if tok, ok := user.Tokens["api"]; ok {
			return tok == token
		}
	}
	return false
}

// Emails returns a list of the emails in the Users map
func (u Users) Emails() []string {
	emails := []string{}
	for email := range u {
		emails = append(emails, email)
	}
	sort.Sort(sort.StringSlice(emails))
	return emails
}

func validate(email string) bool {
	if len(email) < 5 {
		return false
	}
	return true
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}
	return b, nil
}

func generateRandomString(prefix string, s int) (string, error) {
	b, err := generateRandomBytes(s)
	combined := []byte(prefix + ":")
	combined = append(combined, b...)
	return base64.URLEncoding.EncodeToString(combined), err
}

func extractPrefix(token string) (string, error) {
	dec, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return "", err
	}
	if ix := bytes.Index(dec, []byte{':'}); ix >= 0 {
		return string(dec[:ix]), nil
	}
	return string(dec), ErrorInvalidToken
}
