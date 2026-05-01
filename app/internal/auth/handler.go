package auth

import (
	"kilkenny/purpleschool/configs"
	"kilkenny/purpleschool/pkg/req"
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

	mux.HandleFunc("POST /auth/register", hander.register)
	mux.HandleFunc("POST /auth/login", hander.auth)
}

func (hander *AuthHandler) auth(w http.ResponseWriter, r *http.Request) {

	_, err := req.HandleBody[LoginRequest](&w, r)

	if err != nil {
		return
	}

	data := LoginResponse{
		Token: "132123",
	}

	json.Json(w, json.Response{
		Response: data,
		Status:   200,
	})
}

func (hander *AuthHandler) register(w http.ResponseWriter, r *http.Request) {
	_, err := req.HandleBody[RegisterRequest](&w, r)

	if err != nil {
		return
	}

	data := RegisterResponse{
		Token: "132123",
	}

	json.Json(w, json.Response{
		Response: data,
		Status:   200,
	})
}
