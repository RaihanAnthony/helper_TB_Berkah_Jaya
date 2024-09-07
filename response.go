package helper 

import (
	"encoding/json"
	"net/http"
)

func Response(w http.ResponseWriter, payload interface{}, statusCode int) {
	response, _ := json.Marshal(payload)
	w.WriteHeader(statusCode)
	w.Write(response)
}
