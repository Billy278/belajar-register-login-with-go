package service

import (
	"belajar-register-login-with-go/helper"
	"belajar-register-login-with-go/model/domain"
	"belajar-register-login-with-go/model/web"
	"belajar-register-login-with-go/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	Validate       *validator.Validate
	DB             *sql.DB
	RepositoryUser repository.UserRepository
}

func NewUserServiceImpl(validate *validator.Validate, db *sql.DB, repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		Validate:       validate,
		DB:             db,
		RepositoryUser: repository,
	}
}
func (user_service *UserServiceImpl) CreateUser(ctx context.Context, request web.RequestCreate) domain.User {
	err := user_service.Validate.Struct(request)
	helper.PanicIfError(err)
	db := user_service.DB
	user := domain.User{
		Firstname: request.Firstname,
		Lastname:  request.Lastname,
		Username:  request.Username,
		Password:  request.Password,
	}
	user = user_service.RepositoryUser.CreateUser(ctx, db, user)
	return user

}

func (user_service *UserServiceImpl) FindByUsername(ctx context.Context, username string) (domain.User, error) {
	db := user_service.DB
	user, err := user_service.RepositoryUser.FindByUsername(ctx, db, username)
	if err != nil {
		return user, err
	} else {
		return user, nil
	}
}
