package random_api

import (
	"log"
	"net/http"
)

func Start() {
	initServerMux()
}

func initServerMux() {
	mux := http.NewServeMux()
	InitHandlers(mux)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
