package verify

import (
	"kilkenny/purpleschool/configs"
	"net/http"
)

type VerifyHandler struct {
	*configs.Config
}

func VerifyHadlers(mux *http.ServeMux, deps VerifyHandler) {
	handler := VerifyHandler{
		Config: deps.Config,
	}

	mux.HandleFunc("POST /send", handler.send)
	mux.HandleFunc("/verify/{hash}", handler.verify)
}

func (halder *VerifyHandler) send(w http.ResponseWriter, r *http.Request) {
}

func (halder *VerifyHandler) verify(w http.ResponseWriter, r *http.Request) {
}
