package api

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"

	"core/infrastructure"
)

const APIHeader = "X-API-HEADER"

type authMiddleware struct {
	path string
}

func (m *authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	apiKey := r.Header.Get(APIHeader)
	if apiKey == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf(`"missing authentication header %s"`, APIHeader)))
		return
	}

	store, err := infrastructure.NewFileUserStore(m.path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if store.Users.VerifyUser(apiKey) {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func newAuthMiddleware() negroni.Handler {
	return &authMiddleware{path: "users.json"}
}
