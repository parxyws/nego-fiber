package repository

import (
	"context"
	"fmt"

	mtd "github.com/parxyws/nego-fiber/internal/shared/domain"
	"github.com/parxyws/nego-fiber/internal/user"
	"github.com/parxyws/nego-fiber/internal/user/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, entity *domain.User) (*domain.User, error) {
	tx := u.DB.WithContext(ctx)
	err := tx.Transaction(func(tx *gorm.DB) error {
		result := tx.Where("email = ?", entity.Email).Omit("avatar_url", "phone_num").FirstOrCreate(entity)

		if result.RowsAffected == 0 {
			return gorm.ErrRegistered
		}

		if result.Error != nil {
			return fmt.Errorf("UserRepository.CreateUser - %w", result.Error)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *UserRepositoryImpl) UpdateUser(ctx context.Context, entity *domain.User) (*domain.User, error) {
	tx := u.DB.WithContext(ctx)
	err := tx.Transaction(func(tx *gorm.DB) error {
		result := tx.Model(&domain.User{}).Where("id = ?", entity.Id).Updates(entity)

		if result.RowsAffected == 0 {
			return gorm.ErrInvalidData
		}

		if result.Error != nil {
			return fmt.Errorf("UserRepository.CreateUser - %w", result.Error)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *UserRepositoryImpl) DeleteUser(ctx context.Context, entity *domain.User) error {
	tx := u.DB.WithContext(ctx)
	return tx.Transaction(func(tx *gorm.DB) error {
		data := new(domain.User)
		if err := tx.Where("email = ?", entity.Email).First(data).Error; err != nil {
			return fmt.Errorf("UserRepository.DeleteUser - %w", err)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(entity.Password)); err != nil {
			return fmt.Errorf("UserRepository.DeleteUser - %w", err)
		}

		if err := tx.Delete(data).Error; err != nil {
			return fmt.Errorf("UserRepository.DeleteUser - %w", err)
		}

		return nil
	})
}

func (u *UserRepositoryImpl) ReadByUsername(ctx context.Context, entity *domain.User) (*domain.User, error) {
	data := new(domain.User)
	tx := u.DB.WithContext(ctx)

	if err := tx.Where("username = ?", entity.Username).First(data).Error; err != nil {
		return nil, fmt.Errorf("UserRepository.ReadByUsername - %w", err)
	}

	return data, nil
}

func (u *UserRepositoryImpl) ReadByEmail(ctx context.Context, entity *domain.User) (*domain.User, error) {
	data := new(domain.User)
	tx := u.DB.WithContext(ctx)
	if err := tx.Where("email = ?", entity.Email).First(data).Error; err != nil {
		return nil, fmt.Errorf("UserRepository.ReadByEmail - %w", err)
	}

	return data, nil
}

func (u *UserRepositoryImpl) ReadById(ctx context.Context, entity *domain.User) (*domain.User, error) {
	data := new(domain.User)
	tx := u.DB.WithContext(ctx)
	if err := tx.Model(&domain.User{}).Preload("Products").Preload("Auctions").Preload("Orders").Take(data, "id = ?", entity.Id).Error; err != nil {
		return nil, fmt.Errorf("UserRepository.ReadById - %w", err)
	}

	return data, nil
}

func (u *UserRepositoryImpl) ReadAllByRoles(ctx context.Context, metadata *mtd.Metadata) ([]domain.User, error) {
	var users []domain.User
	//tx := u.DB.WithContext(ctx)
	//if err := tx.FindInBatches(); err != nil {
	//	return nil, fmt.Errorf("UserRepository.ReadAllByRoles - %w", err)
	//}

	return users, nil
}
