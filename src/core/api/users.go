package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"core/domain"
	"core/infrastructure"
)

func registerRoutesV1Users(prefix string, r *mux.Router) *mux.Router {
	// unauthenticated routes
	r.HandleFunc(prefix+"{email}", registerUserAction).Methods("POST")

	// authenticated routes
	authenticatedRoutes := mux.NewRouter()
	authenticatedRoutes.HandleFunc(prefix+"", listUsersAction).Methods("GET")
	authenticatedRoutes.NotFoundHandler = NotFoundHandlerJSON()

	n := negroni.New(
		newAuthMiddleware(),
		negroni.Wrap(authenticatedRoutes))
	r.PathPrefix(prefix).Handler(n)

	return r
}

/**
 * actions
 **/

func listUsersAction(w http.ResponseWriter, r *http.Request) {
	users, err := listUsers()
	if err != nil {
		log.Printf("error listing users: %v", err)
		JSONResponse(w, http.StatusInternalServerError, "internal error")
		return
	}
	JSONResponse(w, http.StatusOK, users)
}

func registerUserAction(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	token, err := registerUser(v["email"])
	switch err {
	case nil:
	case domain.ErrorUserExists:
		JSONResponse(w, http.StatusConflict, "conflict")
		return
	default:
		log.Printf("error registering user: %v", err)
		JSONResponse(w, http.StatusInternalServerError, "internal error")
		return
	}
	JSONResponse(w, http.StatusOK, token)
}

/**
 * internal helper functions
 **/

func listUsers() ([]string, error) {
	store, err := infrastructure.NewFileUserStore("users.json")
	if err != nil {
		return nil, err
	}
	return store.Users().Emails(), nil
}

func registerUser(email string) (string, error) {
	store, err := infrastructure.NewFileUserStore("users.json")
	if err != nil {
		return "", err
	}
	token, err := store.Users().RegisterUser(email)
	if err != nil {
		return "", err
	}
	err = store.Save()
	if err != nil {
		return "", err
	}
	return token, nil
}
