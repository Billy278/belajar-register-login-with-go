package helper

import (
	"encoding/json"
	"net/http"
)

func ResponseJson(webResponse interface{}, writer http.ResponseWriter) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusSeeOther)
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	PanicIfError(err)
}

func RuquestBody(requestBody interface{}, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(requestBody)
	PanicIfError(err)
}
