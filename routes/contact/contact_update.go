package contact

import "net/http"

func ContactUpdate(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact Update"))
}
