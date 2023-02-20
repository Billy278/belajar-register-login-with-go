package repository

import (
	"belajar-register-login-with-go/helper"
	"belajar-register-login-with-go/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (user_repository *UserRepositoryImpl) CreateUser(ctx context.Context, db *sql.DB, user domain.User) domain.User {
	sql := "Insert Into user(firstname,lastname,username,password) values(?,?,?,?)"
	result, err := db.ExecContext(ctx, sql, user.Firstname, user.Lastname, user.Username, user.Password)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}

func (user_repository *UserRepositoryImpl) FindByUsername(ctx context.Context, db *sql.DB, username string) (domain.User, error) {
	sql := "Select id,firstname,lastname,username,password from user Where username=?"
	rows, err := db.QueryContext(ctx, sql, username)
	helper.PanicIfError(err)
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Firstname, &user.Lastname, &user.Username, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("NOT FOUND")
	}
}
