package req

import (
	"kilkenny/purpleschool/pkg/res"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {

	body, err := Decode[T](r.Body)

	if err != nil {
		res.Json(*w, res.Response{
			Status:   402,
			Response: err.Error(),
		})

		return nil, err
	}

	err = IsValid[T](body)

	if err != nil {
		res.Json(*w, res.Response{
			Status:   402,
			Response: err.Error(),
		})

		return nil, err
	}

	return &body, nil
}
