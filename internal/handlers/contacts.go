package handlers

import (
	"io"
	"net/http"
)

func contactsHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "contacts\n")
}
