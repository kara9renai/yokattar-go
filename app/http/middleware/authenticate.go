package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/kara9renai/yokattar-go/app/app"
	"github.com/kara9renai/yokattar-go/app/domain/object"
	"github.com/kara9renai/yokattar-go/app/server/handler/httperror"
)

var contextKey struct{}

func Authenticate(app *app.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			a := r.Header.Get("Authentication")
			pair := strings.SplitN(a, " ", 2)
			if len(pair) < 2 {
				httperror.Error(w, http.StatusUnauthorized)
				return
			}

			authType := pair[0]
			if !strings.EqualFold(authType, "username") {
				httperror.Error(w, http.StatusUnauthorized)
				return
			}

			username := pair[1]
			if account, err := app.Dao.Account().FindByUsername(ctx, username); err != nil {
				httperror.InternalServerError(w, err)
				return
			} else if account == nil {
				httperror.Error(w, http.StatusUnauthorized)
				return
			} else {
				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), contextKey, account)))
			}
		})
	}
}

// Read Account data from authorized request
func AccountOf(r *http.Request) *object.Account {
	if cv := r.Context().Value(contextKey); cv == nil {
		return nil
	} else if account, ok := cv.(*object.Account); !ok {
		return nil
	} else {
		return account
	}
}
