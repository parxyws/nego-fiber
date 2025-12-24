package domain

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	auction "github.com/parxyws/nego-fiber/internal/auction/domain"
	order "github.com/parxyws/nego-fiber/internal/order/domain"
	product "github.com/parxyws/nego-fiber/internal/product/domain"
	"gorm.io/gorm"
)

type User struct {
	Id         uuid.UUID         `gorm:"primary_key"`
	Username   string            `gorm:"column:username;not null"`
	Email      string            `gorm:"column:email;not null"`
	Password   string            `gorm:"column:password;not null"`
	FullName   string            `gorm:"column:full_name;not null"`
	PhoneNum   string            `gorm:"column:phone_num;not null"`
	AvatarUrl  string            `gorm:"column:avatar_url"`
	Products   []product.Product `gorm:"foreignKey:user_id;references:id"`
	Orders     []order.Order     `gorm:"foreignKey:user_id;references:id"`
	Auctions   []auction.Auction `gorm:"foreignKey:user_id;references:id"`
	VerifiedAt sql.NullTime      `gorm:"column:verified_at"`
	CreatedAt  time.Time         `gorm:"column:created_at"`
	UpdatedAt  time.Time         `gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt    `gorm:"index;column:deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}
