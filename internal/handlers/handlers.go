package handlers

import (
	"net/http"
)

func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/contacts", contactsHandler)
	return mux
}
