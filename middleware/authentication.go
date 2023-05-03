package middleware

import (
	"net/http"
	"wordfulness/types"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request, *types.User)

type UserStorage interface {
	GetUserByUserName(string) (*types.User, error)
}

func WithAuthentication(
	handler AuthenticatedHandler,
	storage UserStorage,
	loginUrl string,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("authentication")
		if err != nil {
			http.Redirect(w, r, loginUrl, http.StatusMovedPermanently)
			return
		}

		user, err := storage.GetUserByUserName(cookie.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		handler(w, r, user)
	}
}
