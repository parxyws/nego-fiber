package user

import (
	"context"

	"github.com/parxyws/nego-fiber/internal/user/domain/dto"
)

type AuthService interface {
	Register(ctx context.Context, entity *dto.UserRegisterRequest) (*dto.UserRegisterResponse, error)
	Login(ctx context.Context, entity *dto.UserLoginRequest) (*dto.UserResponse, error)
}

type UserService interface {
	ReadCurrentUser(ctx context.Context)
}
