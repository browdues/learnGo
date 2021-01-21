package cancelrequest

import (
	"fmt"
	"net/http"
)

// Store is the store
type Store interface {
	Fetch() string
	Cancel()
}

// Server is a small http server
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}
