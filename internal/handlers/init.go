package handlers

import (
	"net/http"
)

func Initialize() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/contacts", contactsHandler)
}
