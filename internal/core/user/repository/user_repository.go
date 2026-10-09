package repository

import (
	"context"
	"core-banking/internal/core/user/domain"
)

type UserRepository interface {
	Insert(ctx context.Context, user domain.User) (domain.User, error)
}
