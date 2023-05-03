package middleware

import (
	"net/http"
	"strconv"
	"wordfulness/types"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request, *types.User)

type UserStorage interface {
	GetUserById(int) (*types.User, error)
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

		id, err := strconv.ParseInt(cookie.Value, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := storage.GetUserById(int(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		handler(w, r, user)
	}
}
