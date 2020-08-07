package handler

import (
	"fmt"
	"net/http"
)

var count int

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, count)
	count++
}
