package handler

import (
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	http.HandleFunc("/api/count", Handler)
	http.ListenAndServe(":8888", nil)
}
