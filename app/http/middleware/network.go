package middleware

import (
	"net/http"
	"strings"

	"github.com/kara9renai/yokattar-go/app/config"
	"github.com/kara9renai/yokattar-go/app/server/handler/httperror"
)

func AvailableIP() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			nw := config.NewAvaliableNetworks()
			isPrivatedAddr := nw.IsPrivatedAddr(strings.Split(r.RemoteAddr, ":")[0])
			if !isPrivatedAddr {
				httperror.Error(w, http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
