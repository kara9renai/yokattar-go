package network

import (
	"net"
	"net/http"
	"strings"

	"github.com/kara9renai/yokattar-go/app/handler/httperror"
)

type availableNetworks struct {
	networks []*net.IPNet
}

func NewAvaliableNetworks() *availableNetworks {
	nw := new(availableNetworks)
	availableNetWorks := []string{
		// local machine
		"127.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/12",
	}
	for _, v := range availableNetWorks {
		_, ipnet, err := net.ParseCIDR(v)
		if err != nil {
			continue
		}
		nw.networks = append(nw.networks, ipnet)
	}

	return nw
}

func (nw *availableNetworks) IsPrivatedAddr(i string) bool {
	ip := net.ParseIP(i)
	if ip == nil {
		return false
	}
	for _, ipnet := range nw.networks {
		if ipnet.Contains(ip) {
			return true
		}
	}
	return false
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			nw := NewAvaliableNetworks()
			isPrivatedAddr := nw.IsPrivatedAddr(strings.Split(r.RemoteAddr, ":")[0])
			if !isPrivatedAddr {
				httperror.Error(w, http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
