package link

import (
	"kilkenny/purpleschool/pkg/req"
	"kilkenny/purpleschool/pkg/res"
	"net/http"
)

type LinkHandlder struct {
	LinkRepository *LinkRespository
}

type LinkHandlderDeps struct {
	LinkRepository *LinkRespository
}

func LinkHanlder(mux *http.ServeMux, deps LinkHandlderDeps) {
	hand := &LinkHandlder{
		LinkRepository: deps.LinkRepository,
	}

	mux.HandleFunc("GET /{hash}", hand.get)
	mux.HandleFunc("POST /link", hand.create)
	mux.HandleFunc("DELETE /link/{id}", hand.delete)
	mux.HandleFunc("PATCH /link/{id}", hand.update)
}

// GET
func (hand *LinkHandlder) get(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")

	result, err := hand.LinkRepository.Get(hash)

	if err != nil {
		res.Json(w, res.Response{
			Response: err.Error(),
			Status:   http.StatusNotFound,
		})
		return
	}

	http.Redirect(w, r, result.Url, http.StatusTemporaryRedirect)
}

// POSt
func (hand *LinkHandlder) create(w http.ResponseWriter, r *http.Request) {
	body, err := req.HandleBody[LinkCreateRequest](&w, r)

	if err != nil {
		return
	}

	link := NewLink(body.Url)

	createdLink, err := hand.LinkRepository.Create(link)

	if err != nil {
		res.Json(w, res.Response{
			Response: err.Error(),
			Status:   http.StatusInternalServerError,
		})
		return
	}

	res.Json(w, res.Response{
		Response: createdLink,
		Status:   http.StatusCreated,
	})
}

// PATHCH
func (hand *LinkHandlder) update(w http.ResponseWriter, r *http.Request) {}

// DELETE
func (hand *LinkHandlder) delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	hand.LinkRepository.Delete(id)
}
