package middleware

import (
	"net/http"

	"gorm.io/gorm"
)

// HealthCheck middleware
func HealthCheck(db *gorm.DB, dbMedea *gorm.DB, redis interface{}) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// check db connection
			sql, err := db.DB()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = sql.Ping()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			// check db medea connection
			sql, err = dbMedea.DB()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = sql.Ping()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
