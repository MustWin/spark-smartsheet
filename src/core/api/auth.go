package api

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"

	"core/infrastructure"
)

const headerAuthTokenKey = "X-API-TOKEN"

type authMiddleware struct {
	path string
}

func (m *authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	apiKey := r.Header.Get(headerAuthTokenKey)
	if apiKey == "" {
		JSONResponse(w, http.StatusUnauthorized, fmt.Sprintf(`"missing authentication header %s"`, headerAuthTokenKey))
		return
	}

	store, err := infrastructure.NewFileUserStore(m.path)
	if err != nil {
		JSONResponse(w, http.StatusInternalServerError, "internal error")
		return
	}

	if store.Users().VerifyUser(apiKey) {
		next(w, r)
	} else {
		JSONResponse(w, http.StatusUnauthorized, "unauthorized")
		return
	}
}

func newAuthMiddleware() negroni.Handler { return &authMiddleware{path: "users.json"} }
