package application

import (
	"context"
	"errors"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/repository"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/sirupsen/logrus"
)

type addressServiceImpl struct {
	addressRepo repository.AddressRepository
	contactRepo repository.ContactRepository // Needed to check if contact exists
	log         *logrus.Logger
}

func NewAddressService(addressRepo repository.AddressRepository, contactRepo repository.ContactRepository, log *logrus.Logger) service.AddressService {
	return &addressServiceImpl{
		addressRepo: addressRepo,
		contactRepo: contactRepo,
		log:         log,
	}
}

func (s *addressServiceImpl) CreateAddress(ctx context.Context, contactID uint, request service.AddressCreateRequest) (*model.Address, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "CreateAddress", "contactID": contactID})

	// Check if contact exists
	if _, err := s.contactRepo.FindByID(ctx, contactID); err != nil {
		log.WithError(err).Warn("Contact not found for address creation")
		return nil, errors.New("contact not found")
	}

	address := &model.Address{
		ContactID:  contactID,
		Street:     request.Street,
		City:       request.City,
		Province:   request.Province,
		Country:    request.Country,
		PostalCode: request.PostalCode,
	}

	if err := s.addressRepo.Save(ctx, address); err != nil {
		log.WithError(err).Error("Failed to save address")
		return nil, errors.New("failed to create address")
	}

	log.Info("Address created successfully")
	return address, nil
}

func (s *addressServiceImpl) GetAddressByID(ctx context.Context, addressID uint) (*model.Address, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "GetAddressByID", "addressID": addressID})

	address, err := s.addressRepo.FindByID(ctx, addressID)
	if err != nil {
		log.WithError(err).Warn("Address not found")
		return nil, errors.New("address not found")
	}

	log.Info("Successfully retrieved address")
	return address, nil
}

func (s *addressServiceImpl) GetAllAddressesByContactID(ctx context.Context, contactID uint) ([]model.Address, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "GetAllAddressesByContactID", "contactID": contactID})

	addresses, err := s.addressRepo.FindAllByContactID(ctx, contactID)
	if err != nil {
		log.WithError(err).Error("Failed to retrieve addresses for contact")
		return nil, errors.New("server error")
	}

	log.Info("Successfully retrieved all addresses for contact")
	return addresses, nil
}

func (s *addressServiceImpl) UpdateAddress(ctx context.Context, addressID uint, request service.AddressUpdateRequest) (*model.Address, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "UpdateAddress", "addressID": addressID})

	address, err := s.addressRepo.FindByID(ctx, addressID)
	if err != nil {
		log.WithError(err).Warn("Address not found for update")
		return nil, errors.New("address not found")
	}

	address.Street = request.Street
	address.City = request.City
	address.Province = request.Province
	address.Country = request.Country
	address.PostalCode = request.PostalCode

	if err := s.addressRepo.Update(ctx, address); err != nil {
		log.WithError(err).Error("Failed to update address")
		return nil, errors.New("failed to update address")
	}

	log.Info("Address updated successfully")
	return address, nil
}

func (s *addressServiceImpl) DeleteAddress(ctx context.Context, addressID uint) error {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "DeleteAddress", "addressID": addressID})

	if _, err := s.addressRepo.FindByID(ctx, addressID); err != nil {
		log.WithError(err).Warn("Address not found for deletion")
		return errors.New("address not found")
	}

	if err := s.addressRepo.Delete(ctx, addressID); err != nil {
		log.WithError(err).Error("Failed to delete address")
		return errors.New("failed to delete address")
	}

	log.Info("Address deleted successfully")
	return nil
}
