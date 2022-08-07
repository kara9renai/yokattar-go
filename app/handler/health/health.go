package health

import "net/http"

func NewRouter() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		_, err := w.Write([]byte("health check ok"))
		if err != nil {
			panic(err)
		}
	}
}
