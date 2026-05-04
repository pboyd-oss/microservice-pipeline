package handlers

import (
	"encoding/json"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

type validateRequest struct {
	Token string `json:"token"`
}

type validateResponse struct {
	Valid    bool   `json:"valid"`
	Username string `json:"username,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Password == "" {
		http.Error(w, "username and password required", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResponse{Token: "dummy-token-" + req.Username})
}

func Validate(w http.ResponseWriter, r *http.Request) {
	var req validateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if req.Token == "" {
		json.NewEncoder(w).Encode(validateResponse{Valid: false})
		return
	}
	json.NewEncoder(w).Encode(validateResponse{Valid: true, Username: "dummy-user"})
}
