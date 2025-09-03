package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloEndpoint(t *testing.T) {
	var (
		helloEndpoint = "/hello"
	)
	req, err := http.NewRequest(http.MethodGet, helloEndpoint, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello World!"))
	})

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedMsg := "Hello World!"
	if recorder.Body.String() != expectedMsg {
		t.Errorf("handler returned unexpected body: got %v want %v", recorder.Body.String(), expectedMsg)
	}
}
