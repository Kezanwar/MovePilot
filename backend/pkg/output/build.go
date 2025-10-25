package output

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Middleware = func(http.Handler) http.Handler

func MakeSubRouter(r *mux.Router, path string, apply func(*mux.Router)) {
	sr := r.PathPrefix(path).Subrouter()
	apply(sr)
}

func MakeRoute(
	r *mux.Router,
	path string,
	handler JsonHandler,
	middlewares ...Middleware,
) *mux.Route {
	var h http.Handler = MakeJsonHandler(handler)

	// Apply middleware in order
	for _, mw := range middlewares {
		h = mw(h)
	}

	return r.Handle(path, h)
}
