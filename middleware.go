package middleware

import (
	"net/http"

	"github.com/go-martini/martini"
)

type AuthManager interface {
	AuthenticateRequest(req *http.Request) bool
}

func Authentication() martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, amgr AuthManager) {
		if !amgr.AuthenticateRequest(req) {
			http.Error(res, "Not Authorized", http.StatusUnauthorized)
		}
	}
}
