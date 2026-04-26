package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"practice9/idempotency"
	"practice9/retry"
)

func main() {

	fmt.Println("=== RETRY TEST ===")

	counter := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		counter++

		if counter <= 3 {
			w.WriteHeader(503)
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(`{"status":"success"}`))
	}))
	defer server.Close()

	client := retry.Client{
		HttpClient: &http.Client{},
		MaxRetries: 5,
		BaseDelay:  500 * time.Millisecond,
		MaxDelay:   5 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client.ExecutePayment(ctx, server.URL)

	fmt.Println("\n=== IDEMPOTENCY TEST ===")

	store := idempotency.NewStore()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Processing started...")
		time.Sleep(2 * time.Second)
		w.Write([]byte(`{"status":"paid","amount":1000}`))
	})

	server2 := httptest.NewServer(idempotency.Middleware(store, handler))
	defer server2.Close()

	key := "test-key"

	for i := 0; i < 5; i++ {
		go func() {
			req, _ := http.NewRequest("GET", server2.URL, nil)
			req.Header.Set("Idempotency-Key", key)

			resp, _ := http.DefaultClient.Do(req)
			fmt.Println("Status:", resp.StatusCode)
		}()
	}
	time.Sleep(5 * time.Second)

	fmt.Println("\n=== FINAL REQUEST (CACHE CHECK) ===")

	req, _ := http.NewRequest("GET", server2.URL, nil)
	req.Header.Set("Idempotency-Key", key)

	resp, _ := http.DefaultClient.Do(req)
	fmt.Println("Final request status:", resp.StatusCode)
}
