package service

import (
	"context"
	"core-banking/internal/core/user/domain"
	"core-banking/internal/core/user/dto"
)

type UserService interface {
	Profile(ctx context.Context, userID int64) (domain.User, dto.UserResponse, error)
}
