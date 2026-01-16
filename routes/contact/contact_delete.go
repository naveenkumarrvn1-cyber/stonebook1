package contact

import "net/http"

func ContactDelete(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact Delete"))
}
