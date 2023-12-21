package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/adityaqudaedah/go_restful_api/model/web"
)

func WriteToResponseBody(w http.ResponseWriter, response web.WebResponse) {
	w.Header().Set("Content-Type","application/json")

	encoder := json.NewEncoder(w)
	errEncoder := encoder.Encode(response)

	PanicIfError(errEncoder)
}