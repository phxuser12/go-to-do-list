package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func do(h http.Handler, method, path string, form url.Values) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestAddToggleDelete(t *testing.T) {
	store := NewMemoryStore()
	h := NewServer(store).Routes()

	rec := do(h, "POST", "/todos", url.Values{"text": {"buy milk"}})
	if rec.Code != 200 || !strings.Contains(rec.Body.String(), "buy milk") {
		t.Fatalf("add: %d %s", rec.Code, rec.Body)
	}

	do(h, "POST", "/todos/1/toggle", nil)
	if !store.List()[0].Done {
		t.Fatal("expected todo to be done")
	}

	do(h, "DELETE", "/todos/1", nil)
	if len(store.List()) != 0 {
		t.Fatal("expected todo to be deleted")
	}
}

func TestAddIgnoresBlank(t *testing.T) {
	store := NewMemoryStore()
	do(NewServer(store).Routes(), "POST", "/todos", url.Values{"text": {"  "}})
	if len(store.List()) != 0 {
		t.Fatal("blank todo should be ignored")
	}
}
