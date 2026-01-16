package contact

import "net/http"

func ContactList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact List"))
}
