package utils

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SendOk(w http.ResponseWriter) {
	w.Write([]byte("OK"))
}

func SendError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}
