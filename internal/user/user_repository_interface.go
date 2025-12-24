package user

import (
	"context"

	mtd "github.com/parxyws/nego-fiber/internal/shared/domain"
	"github.com/parxyws/nego-fiber/internal/user/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, entity *domain.User) (*domain.User, error)
	UpdateUser(ctx context.Context, entity *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, entity *domain.User) error

	ReadByUsername(ctx context.Context, entity *domain.User) (*domain.User, error)
	ReadByEmail(ctx context.Context, entity *domain.User) (*domain.User, error)
	ReadById(ctx context.Context, entity *domain.User) (*domain.User, error)
	ReadAllByRoles(ctx context.Context, metadata *mtd.Metadata) ([]domain.User, error)
}
