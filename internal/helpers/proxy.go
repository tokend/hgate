package helpers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Proxy func(w http.ResponseWriter, r *http.Request)

func GetProxy(horizonURL *url.URL) Proxy {
	return func(w http.ResponseWriter, r *http.Request) {
		httputil.NewSingleHostReverseProxy(horizonURL).ServeHTTP(w, r)
	}
}
