package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			w.Write([]byte("Hello POST App 2"))
		default:
			w.Write([]byte(fmt.Sprintf("Hello %s App 2", r.Method)))
		}
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		return
	}
}
