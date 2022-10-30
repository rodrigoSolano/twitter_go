package middlew

import (
	"net/http"

	"github.com/rodrigoSolano/twitter_go/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckConnection() {
			http.Error(w, "Connection lost with database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
