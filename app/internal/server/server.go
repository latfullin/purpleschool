package server

import (
	"kilkenny/purpleschool/configs"
	"kilkenny/purpleschool/internal/auth"
	"kilkenny/purpleschool/internal/link"
	"kilkenny/purpleschool/internal/verify"
	"kilkenny/purpleschool/pkg/db"
	"log"
	"net/http"
)

type Repository struct {
	LinkRepository *link.LinkRespository
}

func Start() {
	config := configs.LoadConfig()

	db := db.NewDb(config)

	repository := initRepository(db)

	router := initRouters(config, repository)

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

func initRouters(config *configs.Config, rep *Repository) *http.ServeMux {
	router := http.NewServeMux()
	auth.AuthHandlers(router, auth.AuthHandlerDeps{
		Config: config,
	})

	verify.VerifyHadlers(router, verify.VerifyHandler{
		Config: config,
	})

	link.LinkHanlder(router, link.LinkHandlderDeps{
		LinkRepository: rep.LinkRepository,
	})

	return router
}

func initRepository(db *db.Db) *Repository {
	linkReposotory := link.NewLinkRepository(db)

	return &Repository{
		LinkRepository: linkReposotory,
	}
}
