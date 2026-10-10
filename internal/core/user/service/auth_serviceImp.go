package service

import (
	"context"
	"core-banking/internal/core/user/domain"
	"core-banking/internal/core/user/dto"
	"core-banking/internal/core/user/repository"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImp struct {
	userRepository repository.UserRepository
}

func NewAuthService(
	userRepository repository.UserRepository,
) AuthService {
	return &AuthServiceImp{
		userRepository: userRepository,
	}
}

var (
	// error global yang dipakai setiap kali register gagal (name & password are required)
	ErrInvalidInput      = errors.New("name & password are required")
	ErrUsernameExists    = errors.New("duplikat name")
	ErrInvalidCredential = errors.New("error credencial")
	ErrInternal          = errors.New("internal server error")
)

func (service *AuthServiceImp) Register(ctx context.Context, input dto.RegisterRequest) (dto.AuthResponse, error) {
	if input.Name == "" || input.Password == "" {
		return dto.AuthResponse{}, ErrInvalidInput
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return dto.AuthResponse{}, fmt.Errorf("failed hash :%w", err)
	}

	userDomain := domain.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Status:    "active",
		RoleID:    input.RoleID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	saveUser, err := service.userRepository.Insert(ctx, userDomain)

	if err != nil {
		return dto.AuthResponse{}, err
	}
	userResponse := dto.AuthResponse{
		ID:        saveUser.ID,
		Name:      saveUser.Name,
		Email:     saveUser.Email,
		RoleID:    saveUser.RoleID,
		CreatedAt: saveUser.CreatedAt,
	}
	return userResponse, nil
}
