package handler

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type HelloWorld struct {
	Log *log.Logger
}

func (h *HelloWorld) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
}
