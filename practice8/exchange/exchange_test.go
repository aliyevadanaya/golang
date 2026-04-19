package exchange

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetRateSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"base":"USD","target":"EUR","rate":0.9}`))
	}))
	defer server.Close()

	service := NewExchangeService(server.URL)

	rate, err := service.GetRate("USD", "EUR")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if rate != 0.9 {
		t.Errorf("got %f, want 0.9", rate)
	}
}

func TestGetRateAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":"invalid request"}`))
	}))
	defer server.Close()

	service := NewExchangeService(server.URL)

	_, err := service.GetRate("USD", "EUR")

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestGetRateBadJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`invalid json`))
	}))
	defer server.Close()

	service := NewExchangeService(server.URL)

	_, err := service.GetRate("USD", "EUR")

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestGetRateEmptyBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer server.Close()

	service := NewExchangeService(server.URL)

	_, err := service.GetRate("USD", "EUR")

	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
