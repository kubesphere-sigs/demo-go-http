package handler

import (
	"net/http"
	"testing"
)

func TestHello(t *testing.T) {
	handler := &HelloWorld{}

	req, _ := http.NewRequest(http.MethodGet, "", nil)
	w := &fakeHttpResponseWriter{}

	handler.ServeHTTP(w, req)
}

type fakeHttpResponseWriter struct {
}

func (*fakeHttpResponseWriter) Header() http.Header {
	return nil
}

func (*fakeHttpResponseWriter) Write([]byte) (int, error) {
	return 0, nil
}

func (*fakeHttpResponseWriter) WriteHeader(statusCode int) {
	return
}
