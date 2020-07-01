package web

import (
	"encoding/json"
	"net/http"
)

//RespondJSON :
func RespondJSON(w *http.ResponseWriter, statusCode int, content interface{}) {
	response, err := json.Marshal(content)
	if err != nil {
		writeToHeader(w, http.StatusInternalServerError, err.Error())
		return
	}
	(*w).Header().Set("Content-Type", "application/json")
	writeToHeader(w, statusCode, response)
}

//write to header func
func writeToHeader(w *http.ResponseWriter, statusCode int, payload interface{}) {
	(*w).WriteHeader(statusCode)
	(*w).Write(payload.([]byte))
}
