package handlers

import (
	"io"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
