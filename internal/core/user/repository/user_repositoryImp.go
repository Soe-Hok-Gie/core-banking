package repository

import (
	"context"
	"core-banking/internal/core/user/domain"
	"database/sql"
)

type UserRepositoryImp struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &UserRepositoryImp{
		DB: DB,
	}
}

func (repository *UserRepositoryImp) Insert(ctx context.Context, user domain.User) (domain.User, error) {
	script := "INSERT INTO users (id,name,email,password,status,role_id,created_at, updated,at) VALUES(?,?,?,?,?,?,?,?)"
	result, err := repository.DB.ExecContext(ctx, script, user.ID, user.Name, user.Email, user.Password, user.Status, user.RoleID, user.CreatedAt, user.UpdatedAt)

	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}
	user.ID = id
	return user, nil

}
