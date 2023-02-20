package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Registration(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Home(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	Logout(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
