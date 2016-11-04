package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"

	"core/api"
)

func main() {
	app := cli.NewApp()
	app.Name = "server"
	app.Usage = "start the REST API server"
	app.Version = "0.1.0"
	app.Action = Main
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Main(c *cli.Context) error {
	port := ":8000"

	router := api.RegisterRoutesV1API("/v1/", mux.NewRouter())
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.JSONResponse(w, http.StatusOK, "API at /v1/")
	}).Methods("GET")
	router.NotFoundHandler = api.NotFoundHandlerJSON()

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.UseHandler(router)

	srv := http.Server{
		Handler:      n,
		Addr:         port,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("listening on port %s", port)
	return srv.ListenAndServe()
}
