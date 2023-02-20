package middleware

import (
	"belajar-register-login-with-go/controller"
	"belajar-register-login-with-go/helper"
	"belajar-register-login-with-go/model/web"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type AuthMiddleware struct {
	Handler httprouter.Handle
}

func NewAuthMiddleware(handler httprouter.Handle) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}
func (auth *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	session, _ := controller.Store.Get(request, "Auth-key")
	if len(session.Values) == 0 {
		fmt.Println(session.Values)
		webResponse := web.ResponseWeb{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.ResponseJson(&webResponse, writer)
	} else {
		auth.Handler(writer, request, param)
	}
}
