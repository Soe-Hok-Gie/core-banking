package repository

import (
	"context"
	"core-banking/internal/core/user/domain"
	"database/sql"
	"fmt"
	"log"
)

type UserRepositoryImp struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) UserRepository {
	return &UserRepositoryImp{DB: DB}
}

func (repository *UserRepositoryImp) Insert(ctx context.Context, user domain.User) (domain.User, error) {

	script := "INSERT INTO users (id,name,email,password,status,role_id,created_at, updated_at) VALUES(?,?,?,?,?,?,?,?)"

	// Pastikan repository.DB tidak nil sebelum dieksekusi agar tidak panic
	if repository.DB == nil {
		log.Println("[ERROR] repository.DB bernilai nil! Pastikan koneksi DB dioper saat memanggil NewUserRepository")
		return user, fmt.Errorf("database connection is nil")
	}

	result, err := repository.DB.ExecContext(ctx, script, user.ID, user.Name, user.Email, user.Password, user.Status, user.RoleID, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		// MENCETAK ERROR KE CONSOLE UNTUK DEBUGGING
		log.Println("[DATABASE ERROR] Gagal mengeksekusi query Insert:", err.Error())
		return user, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return user, err
	}
	user.ID = id
	return user, nil

}

func (repository *UserRepositoryImp) GetByEmail(ctx context.Context, email string) (domain.User, error) {

	script := "SELECT id,name,email,password,status,role_id,created_at FROM users WHERE email = ?"
	row := repository.DB.QueryRowContext(ctx, script, email)

	var u domain.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Status, &u.RoleID, &u.CreatedAt)
	if err != nil {
		return u, err
	}
	return u, nil

}

func (repository *UserRepositoryImp) GetById(ctx context.Context, userID int64) (domain.User, error) {
	script := "SELECT id, name, email, password, status, role_id, created_at, updated_at FROM users WHERE id =?"
	row := repository.DB.QueryRowContext(ctx, script, userID)

	var u domain.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Status, &u.RoleID, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return u, err
	}
	return u, nil
}
