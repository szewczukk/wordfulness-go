package middleware

import (
	"context"
	"log"
	"net/http"
	"wordfulness/types"
)

type UserStorage interface {
	GetUserByUserName(string) (*types.User, error)
}

type contextKey int

const userContextKey contextKey = 0

func WithAuthentication(handler http.HandlerFunc, storage UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("authentication")
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusMovedPermanently)
			return
		}

		user, err := storage.GetUserByUserName(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctxWithUser := context.WithValue(r.Context(), userContextKey, user)
		rWithUser := r.WithContext(ctxWithUser)

		handler(w, rWithUser)
	}
}

func UserFromContext(ctx context.Context) (*types.User, bool) {
	user, ok := ctx.Value(userContextKey).(*types.User)
	return user, ok
}
