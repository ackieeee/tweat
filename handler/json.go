package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type JsonHandler func(w http.ResponseWriter, r *http.Request) (int, interface{}, error)

func (h JsonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, payload, err := h(w, r)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	responseJson(w, status, payload)
}

func responseJson(w http.ResponseWriter, status int, payload interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	if err := enc.Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(buf.Bytes()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
