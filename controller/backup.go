package controller

// import (
// 	"belajar-register-login-with-go/helper"
// 	"belajar-register-login-with-go/model/web"
// 	"belajar-register-login-with-go/service"
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/sessions"
// 	"github.com/julienschmidt/httprouter"
// 	"golang.org/x/crypto/bcrypt"
// )

// var key = []byte("Secret-login")
// var store = sessions.NewCookieStore(key)

// type UserControllerImpl struct {
// 	ServiceUser service.UserService
// }

// func NewUserControllerImpl(serviceUser service.UserService) UserController {
// 	return &UserControllerImpl{
// 		ServiceUser: serviceUser,
// 	}

// }
// func (user_controller *UserControllerImpl) Registration(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
// 	if request.Method == http.MethodGet {
// 		webResponse := web.ResponseWeb{
// 			Code:   http.StatusOK,
// 			Status: "Silahkan lakukan Registrasi",
// 		}
// 		helper.ResponseJson(&webResponse, writer)
// 		return
// 	}
// 	requestRegistration := web.RequestCreate{}
// 	helper.RuquestBody(&requestRegistration, request)
// 	_, err := user_controller.ServiceUser.FindByUsername(request.Context(), requestRegistration.Username)

// 	if err == nil {

// 		webResponse := web.ResponseWeb{
// 			Code:   http.StatusOK,
// 			Status: "Username Sudah ada",
// 		}
// 		helper.ResponseJson(&webResponse, writer)
// 		return
// 	}
// 	//default cost = 10
// 	hashPassword, err := bcrypt.GenerateFromPassword([]byte(requestRegistration.Password), bcrypt.DefaultCost)
// 	helper.PanicIfError(err)
// 	requestRegistration.Password = string(hashPassword)
// 	user_controller.ServiceUser.CreateUser(request.Context(), requestRegistration)
// 	webResponse := web.ResponseWeb{
// 		Code:   http.StatusOK,
// 		Status: "Ok",
// 	}
// 	helper.ResponseJson(&webResponse, writer)
// }

// func (user_controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
// 	session, _ := store.Get(request, "Auth-key")
// 	if len(session.Values) != 0 {
// 		fmt.Println(store)
// 		webResponse := web.ResponseWeb{
// 			Code:   http.StatusOK,
// 			Status: "Login Berhasil",
// 		}
// 		helper.ResponseJson(&webResponse, writer)
// 		return
// 	}
// 	if request.Method == http.MethodGet {
// 		webResponse := web.ResponseWeb{
// 			Code:   http.StatusSeeOther,
// 			Status: "Silahkan Masukkan Username dan password",
// 		}
// 		helper.ResponseJson(&webResponse, writer)
// 		return
// 	}
// 	requestLogin := web.RequestLogin{}
// 	helper.RuquestBody(&requestLogin, request)
// 	user, err := user_controller.ServiceUser.FindByUsername(request.Context(), requestLogin.Username)
// 	err2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestLogin.Password))
// 	fmt.Println(err)
// 	fmt.Println(err2)
// 	fmt.Println("TEs====")
// 	if err == nil && err2 == nil {
// 		session, _ := store.Get(request, "Auth-key")
// 		session.Values["username"] = user.Username

// 		err := session.Save(request, writer)
// 		helper.PanicIfError(err)
// 		webResponse := web.ResponseWeb{
// 			Code:   http.StatusSeeOther,
// 			Status: "Login Berhasil",
// 		}
// 		helper.ResponseJson(&webResponse, writer)
// 	} else {
// 		webResponse := web.ResponseWeb{
// 			Code:   http.StatusSeeOther,
// 			Status: "Username atau password salah",
// 		}
// 		helper.ResponseJson(&webResponse, writer)
// 	}

// }

// func (user_controller *UserControllerImpl) Home(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
// 	session, _ := store.Get(request, "Key")
// 	if len(session.Values) == 0 {
// 		webResponse := web.ResponseWeb{
// 			Code:   http.StatusUnauthorized,
// 			Status: "Unauthorized",
// 		}
// 		helper.ResponseJson(&webResponse, writer)
// 	} else {
// 		webResponse := web.ResponseWeb{
// 			Code:   http.StatusOK,
// 			Status: "Selamt datang" + session.Values["username"].(string),
// 			Data:   session.Values["username"].(string),
// 		}
// 		helper.ResponseJson(&webResponse, writer)

// 	}
// }

// func (user_controller *UserControllerImpl) Logout(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
// 	session, _ := store.Get(request, "Auth-key")
// 	session.Options.MaxAge = -1
// 	err := session.Save(request, writer)
// 	helper.PanicIfError(err)
// 	webResponse := web.ResponseWeb{
// 		Code:   http.StatusOK,
// 		Status: "Berhasil Logout",
// 	}
// 	helper.ResponseJson(&webResponse, writer)

// }
