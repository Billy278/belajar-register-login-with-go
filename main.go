package main

import (
	"belajar-register-login-with-go/app"
	"belajar-register-login-with-go/controller"
	"belajar-register-login-with-go/middleware"
	"belajar-register-login-with-go/repository"
	"belajar-register-login-with-go/service"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func main() {
	db := app.NewDB()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logrus.SetOutput(file)
	validate := validator.New()
	userRepository := repository.NewUserRepositoryImpl()
	serviceUser := service.NewUserServiceImpl(validate, db, userRepository)
	controllerUser := controller.NewUserControllerImpl(serviceUser)
	router := httprouter.New()

	router.GET("/login", controllerUser.Login)
	router.POST("/login", controllerUser.Login)
	router.GET("/register", controllerUser.Registration)
	router.POST("/register", controllerUser.Registration)
	router.GET("/Home", middleware.NewAuthMiddleware(controllerUser.Home).ServeHTTP)
	router.GET("/logout", controllerUser.Logout)
	//router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:9000",
		Handler: router,
	}
	server.ListenAndServe()

}
