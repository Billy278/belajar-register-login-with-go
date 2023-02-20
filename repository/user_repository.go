package repository

import (
	"belajar-register-login-with-go/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, db *sql.DB, user domain.User) domain.User
	FindByUsername(ctx context.Context, db *sql.DB, username string) (domain.User, error)
}
