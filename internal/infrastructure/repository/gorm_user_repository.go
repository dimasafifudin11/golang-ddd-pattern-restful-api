package repository

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/repository"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) Save(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *gormUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *gormUserRepository) FindByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
