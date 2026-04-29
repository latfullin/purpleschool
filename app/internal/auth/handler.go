package auth

import (
	"fmt"
	"kilkenny/purpleschool/configs"
	json "kilkenny/purpleschool/pkg/res"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func AuthHandlers(mux *http.ServeMux, deps AuthHandlerDeps) {
	hander := &AuthHandler{
		Config: deps.Config,
	}

	mux.HandleFunc("/auth/register", hander.register)
	mux.HandleFunc("/auth/login", hander.auth)
}

func (hander *AuthHandler) auth(w http.ResponseWriter, r *http.Request) {
	fmt.Println(hander.Server.Addr)

	data := LoginResponse{
		Token: "132123",
	}

	json.Json(w, json.Response{
		Response: data,
		Status:   200,
	})
}

func (hander *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
}
