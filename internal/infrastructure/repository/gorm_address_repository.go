package repository

import (
	"context"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/repository"
	"gorm.io/gorm"
)

type gormAddressRepository struct {
	db *gorm.DB
}

func NewGormAddressRepository(db *gorm.DB) repository.AddressRepository {
	return &gormAddressRepository{db: db}
}

func (r *gormAddressRepository) Save(ctx context.Context, address *model.Address) error {
	return r.db.WithContext(ctx).Create(address).Error
}

func (r *gormAddressRepository) FindByID(ctx context.Context, addressID uint) (*model.Address, error) {
	var address model.Address
	err := r.db.WithContext(ctx).First(&address, addressID).Error
	return &address, err
}

func (r *gormAddressRepository) FindAllByContactID(ctx context.Context, contactID uint) ([]model.Address, error) {
	var addresses []model.Address
	err := r.db.WithContext(ctx).Where("contact_id = ?", contactID).Find(&addresses).Error
	return addresses, err
}

func (r *gormAddressRepository) Update(ctx context.Context, address *model.Address) error {
	return r.db.WithContext(ctx).Save(address).Error
}

func (r *gormAddressRepository) Delete(ctx context.Context, addressID uint) error {
	return r.db.WithContext(ctx).Delete(&model.Address{}, addressID).Error
}
