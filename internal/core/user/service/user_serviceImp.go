package service

import (
	"context"
	"core-banking/internal/core/user/domain"
	"core-banking/internal/core/user/dto"
	"core-banking/internal/core/user/repository"
)

type userServiceImp struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImp) Profile(ctx context.Context, userID int64) (domain.User, dto.UserResponse, error) {
	//ambil data asli dari domain
	userDomain, err := service.UserRepository.profile(ctx, userID)
	if err != nil {
		return domain.User{}, dto.UserResponse{}, err
	}

	//konversi RoleID(int64) pada domain menjadi role (string) pada DTO
	var roleName string
	switch userDomain.RoleID {
	case 1:
		roleName = "admin"
	case 2:
		roleName = "nasabah prioritas"
	case 3:
		roleName = "nasabah"
	default:
		roleName = "unknown"
	}

	//msping data domain ke dto (field password jangan)
	userDto := dto.UserResponse{
		ID:    userDomain.ID,
		Name:  userDomain.Name,
		Email: userDomain.Email,
		Role:  roleName,
	}
	return userDomain, userDto, err
}
