package request

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/kara9renai/yokattar-go/pkg/config"
	"github.com/pkg/errors"
)

func IDOf(r *http.Request) (int64, error) {
	ids := chi.URLParam(r, "id")

	if ids == "" {
		return -1, errors.Errorf("id was not presenced")
	}

	id, err := strconv.ParseInt(ids, 10, 64)

	if err != nil {
		return -1, errors.Errorf("id was not number")
	}

	return id, nil
}

func URLParamOf(r *http.Request, val string) (int64, error) {
	a, err := strconv.Atoi(r.URL.Query().Get(val))

	if err != nil {
		return 0, err
	}

	return int64(a), nil
}

func UsernameOf(r *http.Request) string {
	return chi.URLParam(r, "username")
}

// TODO: utilsクラスを新たに作った方が良いかも
func LimitOf(r *http.Request) int64 {
	limit, err := URLParamOf(r, "limit")
	if err != nil {
		limit = config.DEFAULT_LIMIT
	}
	if limit > config.MAX_LIMIT {
		limit = config.MAX_LIMIT
	}
	return limit
}
