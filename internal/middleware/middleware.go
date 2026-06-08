package middleware

import "net/http"

type MiddlewareFunc func(http.Handler) http.Handler

func CreateMiddlewareStack(handler http.Handler, m ...MiddlewareFunc) http.Handler {

	for i := len(m) - 1; i > 0; i-- {
		handler = m[i](handler)
	}
	return handler

}
