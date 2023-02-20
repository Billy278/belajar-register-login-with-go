package controller

import (
	"belajar-register-login-with-go/helper"
	"belajar-register-login-with-go/model/web"
	"belajar-register-login-with-go/service"
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var key = []byte("Secret-login")
var Store = sessions.NewCookieStore(key)

type UserControllerImpl struct {
	ServiceUser service.UserService
}

func NewUserControllerImpl(serviceUser service.UserService) UserController {
	return &UserControllerImpl{
		ServiceUser: serviceUser,
	}

}
func (user_controller *UserControllerImpl) Registration(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	if request.Method == http.MethodGet {
		webResponse := web.ResponseWeb{
			Code:   http.StatusOK,
			Status: "Silahkan lakukan Registrasi",
		}
		helper.ResponseJson(&webResponse, writer)
		return
	}
	requestRegistration := web.RequestCreate{}
	fmt.Println(requestRegistration.Username)
	helper.RuquestBody(&requestRegistration, request)
	_, err := user_controller.ServiceUser.FindByUsername(request.Context(), requestRegistration.Username)

	if err == nil {

		webResponse := web.ResponseWeb{
			Code:   http.StatusOK,
			Status: "Username Sudah ada",
		}
		helper.ResponseJson(&webResponse, writer)
		return
	}
	//default cost = 10
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(requestRegistration.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	requestRegistration.Password = string(hashPassword)
	user_controller.ServiceUser.CreateUser(request.Context(), requestRegistration)
	webResponse := web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Ok",
	}
	logrus.WithFields(logrus.Fields{
		"newUser": requestRegistration.Username,
	}).Info("Berhasil Registrasi")
	helper.ResponseJson(&webResponse, writer)
}

func (user_controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	session, _ := Store.Get(request, "Auth-key")
	if len(session.Values) != 0 {
		fmt.Println(Store)
		http.Redirect(writer, request, "/home", http.StatusSeeOther)
		return
	}
	if request.Method == http.MethodGet {
		webResponse := web.ResponseWeb{
			Code:   http.StatusSeeOther,
			Status: "Silahkan Masukkan Username dan password",
		}
		helper.ResponseJson(&webResponse, writer)
		return
	}
	requestLogin := web.RequestLogin{}
	helper.RuquestBody(&requestLogin, request)
	user, err := user_controller.ServiceUser.FindByUsername(request.Context(), requestLogin.Username)
	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestLogin.Password))

	if err == nil && err2 == nil {
		session, _ := Store.Get(request, "Auth-key")
		session.Values["username"] = user.Username

		err := session.Save(request, writer)
		helper.PanicIfError(err)
		logrus.WithFields(logrus.Fields{
			"username": session.Values["username"].(string),
		}).Info("Login berhasil")

		http.Redirect(writer, request, "/home", http.StatusSeeOther)
	} else {
		webResponse := web.ResponseWeb{
			Code:   http.StatusSeeOther,
			Status: "Username atau password salah",
		}
		helper.ResponseJson(&webResponse, writer)
	}

}

func (user_controller *UserControllerImpl) Home(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	session, _ := Store.Get(request, "Auth-key")
	logrus.WithFields(logrus.Fields{
		"username": session.Values["username"].(string),
	}).Info("selamat datang di menu Home")
	webResponse := web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Selamat datang " + session.Values["username"].(string),
		Data:   session.Values["username"].(string),
	}
	helper.ResponseJson(&webResponse, writer)

}

func (user_controller *UserControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	session, _ := Store.Get(request, "Auth-key")
	if session.Values["username"] != nil {
		logrus.WithFields(logrus.Fields{
			"username": session.Values["username"].(string),
		}).Info("Berhasil Logout")
	}

	session.Options.MaxAge = -1
	err := session.Save(request, writer)
	helper.PanicIfError(err)
	webResponse := web.ResponseWeb{
		Code:   http.StatusOK,
		Status: "Berhasil Logout",
	}
	helper.ResponseJson(&webResponse, writer)

}
