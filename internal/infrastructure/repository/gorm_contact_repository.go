package repository

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/repository"
	"gorm.io/gorm"
)

type gormContactRepository struct {
	db *gorm.DB
}

func NewGormContactRepository(db *gorm.DB) repository.ContactRepository {
	return &gormContactRepository{db: db}
}

func (r *gormContactRepository) Save(ctx context.Context, contact *model.Contact) error {
	return r.db.WithContext(ctx).Create(contact).Error
}

func (r *gormContactRepository) FindByID(ctx context.Context, contactID uint) (*model.Contact, error) {
	var contact model.Contact
	err := r.db.WithContext(ctx).Preload("Addresses").First(&contact, contactID).Error
	return &contact, err
}

func (r *gormContactRepository) FindAllByUserID(ctx context.Context, userID uint) ([]model.Contact, error) {
	var contacts []model.Contact
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&contacts).Error
	return contacts, err
}

func (r *gormContactRepository) Update(ctx context.Context, contact *model.Contact) error {
	return r.db.WithContext(ctx).Save(contact).Error
}

func (r *gormContactRepository) Delete(ctx context.Context, contactID uint) error {
	return r.db.WithContext(ctx).Delete(&model.Contact{}, contactID).Error
}
