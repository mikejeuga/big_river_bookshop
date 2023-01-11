package auth

import "net/http"

type Config struct {
	BigSecret string `envconfig:"BIG_SECRET"`
}

func FOMW(theBigSecret string) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			headerVal := r.Header.Get("X-API-KEY")
			if headerVal != theBigSecret {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
