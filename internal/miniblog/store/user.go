package store

import (
	"context"

	"github.com/cjhw/miniblog/internal/pkg/model"
	"gorm.io/gorm"
)

// UserStore 定义了 user 模块在 store 层所实现的方法.
type UserStore interface {
	Create(ctx context.Context, user *model.UserM) error
}

// UserStore 接口的实现.
type users struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ UserStore = (*users)(nil)

func NewUsers(db *gorm.DB) *users {
	return &users{db}
}

func (u *users) Create(ctx context.Context, user *model.UserM) error {
	return u.db.Create(&user).Error
}
