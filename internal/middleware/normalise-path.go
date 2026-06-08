package middleware

import (
	"net/http"
	"strings"
)

func NormalizePath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Keep "/" as-is.
		if path != "/" && strings.HasSuffix(path, "/") {
			path = strings.TrimSuffix(path, "/")

			if r.URL.RawQuery != "" {
				path += "?" + r.URL.RawQuery
			}

			http.Redirect(w, r, path, http.StatusMovedPermanently)
			return
		}

		next.ServeHTTP(w, r)
	})
}
