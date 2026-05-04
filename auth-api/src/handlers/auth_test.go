package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	body, _ := json.Marshal(loginRequest{Username: "alice", Password: "secret"})
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	w := httptest.NewRecorder()
	Login(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp tokenResponse
	json.NewDecoder(w.Body).Decode(&resp)
	if resp.Token == "" {
		t.Fatal("expected non-empty token")
	}
}

func TestLoginMissingFields(t *testing.T) {
	body, _ := json.Marshal(loginRequest{})
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewReader(body))
	w := httptest.NewRecorder()
	Login(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestValidate(t *testing.T) {
	body, _ := json.Marshal(validateRequest{Token: "some-token"})
	req := httptest.NewRequest(http.MethodPost, "/auth/validate", bytes.NewReader(body))
	w := httptest.NewRecorder()
	Validate(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var resp validateResponse
	json.NewDecoder(w.Body).Decode(&resp)
	if !resp.Valid {
		t.Fatal("expected valid=true")
	}
}
