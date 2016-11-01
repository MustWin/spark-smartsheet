package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"core/infrastructure"
)

// RegisterV1UsersRoutes ensures that the routes that handle user
// related actions are registered
func RegisterV1UsersRoutes(prefix string, r *mux.Router) {
	// unauthenticated
	r.HandleFunc(prefix+"{email}", registerUserAction).Methods("POST")

	// authenticated
	authenticatedRoutes := mux.NewRouter()
	authenticatedRoutes.HandleFunc(prefix+"", listUsersAction)

	n := negroni.New(
		newAuthMiddleware(),
		negroni.Wrap(authenticatedRoutes))

	r.PathPrefix(prefix).Handler(n)
}

func listUsersAction(w http.ResponseWriter, r *http.Request) {
	users, err := listUsers()
	if err != nil {
		log.Printf("error listing users: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf(`[\n"%s"\n]`, strings.Join(users, `",\n"`))))
}

func registerUserAction(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	token, err := registerUser(v["email"])
	if err != nil {
		log.Printf("error registering user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf(`%q`, token)))
}

func listUsers() ([]string, error) {
	store, err := infrastructure.NewFileUserStore("users.json")
	if err != nil {
		return nil, err
	}
	users := make([]string, len(store.Users))
	ix := 0
	for k, _ := range store.Users {
		users[ix] = k
		ix++
	}
	return users, nil
}

func registerUser(email string) (string, error) {
	store, err := infrastructure.NewFileUserStore("users.json")
	if err != nil {
		return "", err
	}
	token, err := store.RegisterUser(email)
	if err != nil {
		return "", err
	}
	err = store.Save()
	if err != nil {
		return "", err
	}
	return token, nil
}
