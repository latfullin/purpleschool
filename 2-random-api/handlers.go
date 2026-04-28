package random_api

import (
	"fmt"
	"math/rand"
	"net/http"
)

func InitHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/random-number", handleRandomNumber)
}

func handleRandomNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rand := rand.Intn(6) + 1

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%d", rand)
}
