package res

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status   int
	Response any
}

func Json(w http.ResponseWriter, r Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r.Response)
}
