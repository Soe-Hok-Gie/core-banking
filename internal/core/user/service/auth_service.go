package service

import (
	"context"
	"core-banking/internal/core/user/dto"
)

type AuthService interface {
	Register(ctx context.Context, input dto.RegisterRequest) (dto.AuthResponse, error)
}
