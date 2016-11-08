package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func registerRoutesV1Hooks(prefix string, r *mux.Router) *mux.Router {
	r.HandleFunc(prefix+"{email}", defaultHookAction).Methods("POST")
	return r
}

func defaultHookAction(w http.ResponseWriter, r *http.Request) {
	JSONResponse(w, http.StatusOK, "ok")
}
