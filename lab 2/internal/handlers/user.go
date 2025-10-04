package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("id")
	if q == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}
	id, err := strconv.Atoi(q)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid id"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]int{"user_id": id})
}

type createReq struct {
	Name string `json:"name"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req createReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid name"})
		return
	}
	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid name"})
		return
	}
	writeJSON(w, http.StatusCreated, map[string]string{"created": req.Name})
}
