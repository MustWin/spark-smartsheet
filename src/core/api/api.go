package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
)

type registerRouteFn func(string, *mux.Router) *mux.Router

// RegisterRoutesV1API registers the routes for v1 of the API
func RegisterRoutesV1API(prefix string, r *mux.Router) *mux.Router {
	// v1 routes
	routes := map[string]registerRouteFn{
		prefix + "users/": registerRoutesV1Users,
		prefix + "hooks/": registerRoutesV1Hooks,
	}

	router := r
	for routePrefix, routeFn := range routes {
		router = routeFn(routePrefix, router)
	}

	endpoints, err := describeRoutes(r)
	if err != nil {
		log.Printf("error walking routes: %v", err)
	}
	sort.Sort(sort.StringSlice(endpoints))

	defaultAction := func(w http.ResponseWriter, r *http.Request) {
		resp := struct {
			Description string   `json:"description"`
			Endpoints   []string `json:"endpoints"`
		}{Description: "API v1 Endpoints", Endpoints: endpoints}
		JSONResponse(w, http.StatusOK, resp)
	}

	r.HandleFunc(prefix, defaultAction).Methods("GET")
	r.NotFoundHandler = NotFoundHandlerJSON()
	return r
}

// NotFoundHandlerJSON serializes the not found response as JSON
func NotFoundHandlerJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		JSONResponse(w, http.StatusNotFound, "404 not found")
	}
}

// JSONResponse sends the obj as a response that is JSON encoded
func JSONResponse(w http.ResponseWriter, status int, obj interface{}) {
	w.Header()["Content-Type"] = []string{"application/json; charset=utf8"}
	w.WriteHeader(status)
	jsonTo(w, obj)
}

func jsonTo(w io.Writer, obj interface{}) {
	msg, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Printf("error marshalling response: %v", err)
		msg = []byte(`"error displaying response"`)
	}
	if _, err = w.Write(msg); err != nil {
		log.Printf("error writing response: %v", err)
	}
}

func describeRoutes(r *mux.Router) ([]string, error) {
	routes := []string{}
	return routes, r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		templ, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		routes = append(routes, templ)
		return nil
	})
}
