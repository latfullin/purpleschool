package verify

import (
	"fmt"
	"kilkenny/purpleschool/configs"
	"kilkenny/purpleschool/internal/mailing"
	"kilkenny/purpleschool/pkg/helpers"
	"kilkenny/purpleschool/pkg/req"
	"kilkenny/purpleschool/pkg/res"
	"net/http"

	"github.com/go-playground/validator/v10"
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
	body, err := req.HandleBody[SendRequest](&w, r)

	if err != nil {
		return
	}

	hash, err := helpers.GenerateHash()
	_, fail := helpers.SaveHash(hash)

	if err != nil || fail != nil {
		res.Json(w, res.Response{
			Response: "Ошибка при подготовке данных.",
			Status:   http.StatusInternalServerError,
		})
	}

	if err := mailing.Send(&mailing.Sender{
		To:      body.Email,
		Name:    "example",
		Subject: "Link auth",
		HTML:    fmt.Sprintf("<h1>Подтвердить учетную запись</h1><p><a href='http://localhost:8080/verify/%s'>Подтвердить</a></p>", hash),
	}); err != nil {
		res.Json(w, res.Response{
			Response: "Ошибка при отправке письма.",
			Status:   http.StatusInternalServerError,
		})
		return
	}

	res.Json(w, res.Response{
		Response: "Письмо отправлено на почту.",
		Status:   http.StatusOK,
	})
}

func (halder *VerifyHandler) verify(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")
	hashWr, errRead := helpers.ReadHash()

	if errRead != nil {
		res.Json(w, res.Response{
			Response: false,
			Status:   http.StatusBadRequest,
		})
		return
	}

	payload := ConfirmPayload{
		Hash: hash,
	}

	if err := validator.New().Struct(payload); err != nil {
		res.Json(w, res.Response{
			Response: err.Error(),
			Status:   http.StatusBadRequest,
		})
		return
	}

	if hash == "" {
		res.Json(w, res.Response{
			Response: "Не передан hash.",
			Status:   http.StatusBadRequest,
		})
		return
	}

	if hash == hashWr {
		res.Json(w, res.Response{
			Response: true,
			Status:   http.StatusOK,
		})
	} else {
		helpers.DelereHash()
		res.Json(w, res.Response{
			Response: false,
			Status:   http.StatusBadRequest,
		})
	}
}
