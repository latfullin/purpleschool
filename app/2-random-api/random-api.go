package random_api

import (
	"kilkenny/purpleschool/configs"
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
		Addr:    configs.LoadConfig().Server.Addr,
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}
