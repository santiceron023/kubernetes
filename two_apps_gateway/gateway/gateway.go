package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func proxyRequest(target string, w http.ResponseWriter, r *http.Request) {
	targetURL, err := url.Parse(target)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	proxy.ServeHTTP(w, r)
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/app1" || r.URL.Path == "/app1/":
		proxyRequest("http://app1:8888", w, r)
	case r.URL.Path == "/app2" || r.URL.Path == "/app2/":
		proxyRequest("http://app2:8888", w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handleRequests)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
