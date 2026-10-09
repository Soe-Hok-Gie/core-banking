package service

import (
	"context"
	"core-banking/internal/core/user/dto"
	"errors"
)

type AuthServiceImp struct {
}

var (
	// error global yang dipakai setiap kali register gagal (name & password are required)
	ErrInvalidInput      = errors.New("name & password are required")
	ErrUsernameExists    = errors.New("duplikat name")
	ErrInvalidCredential = errors.New("error credencial")
	ErrInternal          = errors.New("internal server error")
)

func (service *AuthServiceImp) Register(ctx context.Context, input dto.RegisterRequest) (dto.AuthResponse, error) {

}
