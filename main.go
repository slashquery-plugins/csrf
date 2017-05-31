package csrf

import (
	"log"
	"net/http"
	"time"
)

func CSRF(next http.Handler) http.Handler {
	time := time.Now().UTC().String()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("CSRF init")
		w.Header().Set("sq-CSRF-version", time)
		r.Header.Set("sq-CSRF-version", time)
		next.ServeHTTP(w, r)
	})
}
