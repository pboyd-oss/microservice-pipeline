package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestList(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/data", nil)
	w := httptest.NewRecorder()
	List(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var items []dataItem
	json.NewDecoder(w.Body).Decode(&items)
	if len(items) == 0 {
		t.Fatal("expected non-empty item list")
	}
}

func TestGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/data/1", nil)
	req.SetPathValue("id", "1")
	w := httptest.NewRecorder()
	Get(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var item dataItem
	json.NewDecoder(w.Body).Decode(&item)
	if item.ID != "1" {
		t.Fatalf("expected id 1, got %s", item.ID)
	}
}

func TestGetNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/data/999", nil)
	req.SetPathValue("id", "999")
	w := httptest.NewRecorder()
	Get(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}
