package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterV1API(prefix string, r *mux.Router) {
	RegisterV1UsersRoutes(prefix+"users/", r)

	r.HandleFunc(prefix+"", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`"v1 of the API"`))
	})
}
