package main

import "net/http"

func main() {
	SetRouting()
	http.ListenAndServe(":3100", nil)
}

func SetRouting() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "ok"}`))
	})
}
