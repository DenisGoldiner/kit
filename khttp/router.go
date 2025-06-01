package khttp

import (
	"fmt"
	"net/http"
)

// NewRouter prepares the HTTP router.
func NewRouter(baseURL string, handlers map[string]http.Handler) http.Handler {
	mux := http.NewServeMux()

	for route, handler := range handlers {
		mux.Handle(buildHandlerURL(baseURL, route), handler)
	}

	return mux
}

// buildHandlerURL adds the appBaseURL prefix to handler url.
func buildHandlerURL(baseURL, route string) string {
	return fmt.Sprintf("%s%s", baseURL, route)
}
