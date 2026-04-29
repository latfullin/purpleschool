package server

import (
	"kilkenny/purpleschool/configs"
	"kilkenny/purpleschool/internal/auth"
	"kilkenny/purpleschool/internal/verify"
	"log"
	"net/http"
)

func Start() {
	config := configs.LoadConfig()
	router := initRouters(config)
	initServerMux(router, config)
}

func initServerMux(mux *http.ServeMux, config *configs.Config) {
	server := http.Server{
		Addr:    config.Server.Addr,
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}
}

func initRouters(config *configs.Config) *http.ServeMux {
	router := http.NewServeMux()
	auth.AuthHandlers(router, auth.AuthHandlerDeps{
		Config: config,
	})

	verify.VerifyHadlers(router, verify.VerifyHandler{
		Config: config,
	})

	return router
}
