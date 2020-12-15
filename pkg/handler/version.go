package handler

import (
	"fmt"
	"log"
	"net/http"
)

type Version struct {
	Log *log.Logger
}

func (h *Version) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version 3")
}
