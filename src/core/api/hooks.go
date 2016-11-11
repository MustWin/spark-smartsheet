package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func registerRoutesV1Hooks(prefix string, r *mux.Router) *mux.Router {
	r.HandleFunc(prefix+"spark/{email}", sparkCallbackAction).Methods("POST")
	return r
}

func sparkCallbackAction(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)

	JSONResponse(w, http.StatusOK, "ok "+v["email"])
}
