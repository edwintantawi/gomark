package helper

import (
	"encoding/json"
	"github.com/edwintantawi/gomark/src/domain/response"
	"net/http"
)

func ResponseJson(w http.ResponseWriter, code int, body response.Body) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	jsonResponse, err := json.Marshal(body)
	PanicError(err)

	_, err = w.Write(jsonResponse)
	PanicError(err)
}
