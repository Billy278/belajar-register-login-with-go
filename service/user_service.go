package service

import (
	"belajar-register-login-with-go/model/domain"
	"belajar-register-login-with-go/model/web"
	"context"
)

type UserService interface {
	CreateUser(ctx context.Context, request web.RequestCreate) domain.User
	FindByUsername(ctx context.Context, username string) (domain.User, error)
}
