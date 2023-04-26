package main

import "net/http"

func main() {
	SetRouting()
	http.ListenAndServe(":8080", nil)
}

func SetRouting() {
	http.Handle("/", setCacheControl(http.FileServer(http.Dir("dist"))))

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "ok"}`))
	})
}

// cache 設定
func setCacheControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=86400")
		h.ServeHTTP(w, r)
	})
}
