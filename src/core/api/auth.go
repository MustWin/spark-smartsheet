package api

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"

	"core/infrastructure"
)

const headerAuthTokenKey = "Authorization"
const bearerPrefix = "Bearer "

type authMiddleware struct {
	path string
}

func (m *authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	apiKey := r.Header.Get(headerAuthTokenKey)
	if apiKey == "" {
		JSONResponse(w, http.StatusUnauthorized, fmt.Sprintf(`missing authorization header "%s: %s<api token>"`, headerAuthTokenKey, bearerPrefix))
		return
	}

	if len(apiKey) <= len(bearerPrefix) || apiKey[:len(bearerPrefix)] != bearerPrefix {
		JSONResponse(w, http.StatusUnauthorized, fmt.Sprintf(`missing bearer api token in  header "%s: %s<api token>"`, headerAuthTokenKey, bearerPrefix))
		return
	}

	store, err := infrastructure.NewFileUserStore(m.path)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	if store.Users().VerifyUser(apiKey[len(bearerPrefix):]) {
		next(w, r)
	} else {
		JSONResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}
}

func newAuthMiddleware() negroni.Handler { return &authMiddleware{path: "users.json"} }
