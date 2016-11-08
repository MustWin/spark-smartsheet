package api

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"

	"github.com/urfave/negroni"
)

// traceMiddleware provides a logging middleware for tracing requests.
// https://justinas.org/writing-http-middleware-in-go/ for inspiration.
type traceMiddleware struct {
	trace bool
}

func (m *traceMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if m.trace {
		m.debug("REQUEST")(httputil.DumpRequest(r, true))

		rec := httptest.NewRecorder()

		next(rec, r)

		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(rec.Code)
		w.Write(rec.Body.Bytes())

		m.debug("RESPONSE")(httputil.DumpResponse(rec.Result(), true))
		return
	}
	next(w, r)
}

func (m *traceMiddleware) debug(label string) func([]byte, error) {
	wrapper := "\n" +
		"------------------\n" +
		"     " + label + "\n" +
		"------------------\n%s\n" +
		"------------------\n"
	return func(data []byte, err error) {
		if err == nil {
			log.Printf(wrapper, string(data))
		} else {
			log.Printf("\n%s\nERROR: %v\n", label, err)
		}
	}
}

func NewTraceMiddleware(enabled bool) negroni.Handler {
	return &traceMiddleware{trace: enabled}
}
