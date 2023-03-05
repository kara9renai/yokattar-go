package like

import "net/http"

func (h *handler) Like(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
