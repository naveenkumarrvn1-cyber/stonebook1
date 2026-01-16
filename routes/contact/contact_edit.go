package contact

import "net/http"

func ContactEdit(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Contact Edit"))
}
