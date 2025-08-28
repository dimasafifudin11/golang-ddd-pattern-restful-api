package application

import (
	"context"
	"errors"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/common"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/repository"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type contactServiceImpl struct {
	contactRepo repository.ContactRepository
	userRepo    repository.UserRepository // Needed to check if user exists
	log         *logrus.Logger
}

func NewContactService(contactRepo repository.ContactRepository, userRepo repository.UserRepository, log *logrus.Logger) service.ContactService {
	return &contactServiceImpl{
		contactRepo: contactRepo,
		userRepo:    userRepo,
		log:         log,
	}
}

func (s *contactServiceImpl) CreateContact(ctx context.Context, userID uint, request service.ContactCreateRequest) (*model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "CreateContact", "userID": userID})

	_, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.WithError(err).Warn("User not found for contact creation")
			return nil, common.ErrNotFound // <-- User tidak ditemukan
		}
		return nil, errors.New("server error")
	}

	contact := &model.Contact{
		UserID:    userID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Phone:     request.Phone,
	}

	if err := s.contactRepo.Save(ctx, contact); err != nil {
		log.WithError(err).Error("Failed to save contact")
		return nil, errors.New("failed to create contact")
	}

	log.Info("Contact created successfully")
	return contact, nil
}

func (s *contactServiceImpl) GetContactByID(ctx context.Context, contactID uint) (*model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "GetContactByID", "contactID": contactID})

	contact, err := s.contactRepo.FindByID(ctx, contactID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.WithError(err).Warn("Contact not found")
			return nil, common.ErrNotFound // <-- Contact tidak ditemukan
		}
		return nil, errors.New("server error")
	}

	log.Info("Successfully retrieved contact")
	return contact, nil
}

func (s *contactServiceImpl) GetAllContactsByUserID(ctx context.Context, userID uint) ([]model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "GetAllContactsByUserID", "userID": userID})

	contacts, err := s.contactRepo.FindAllByUserID(ctx, userID)
	if err != nil {
		log.WithError(err).Error("Failed to retrieve contacts for user")
		return nil, errors.New("server error")
	}

	log.Info("Successfully retrieved all contacts for user")
	return contacts, nil
}

func (s *contactServiceImpl) UpdateContact(ctx context.Context, contactID uint, request service.ContactUpdateRequest) (*model.Contact, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "UpdateContact", "contactID": contactID})

	contact, err := s.contactRepo.FindByID(ctx, contactID)
	if err != nil {
		log.WithError(err).Warn("Contact not found for update")
		return nil, errors.New("contact not found")
	}

	contact.FirstName = request.FirstName
	contact.LastName = request.LastName
	contact.Email = request.Email
	contact.Phone = request.Phone

	if err := s.contactRepo.Update(ctx, contact); err != nil {
		log.WithError(err).Error("Failed to update contact")
		return nil, errors.New("failed to update contact")
	}

	log.Info("Contact updated successfully")
	return contact, nil
}

func (s *contactServiceImpl) DeleteContact(ctx context.Context, contactID uint) error {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "DeleteContact", "contactID": contactID})

	if _, err := s.contactRepo.FindByID(ctx, contactID); err != nil {
		log.WithError(err).Warn("Contact not found for deletion")
		return errors.New("contact not found")
	}

	if err := s.contactRepo.Delete(ctx, contactID); err != nil {
		log.WithError(err).Error("Failed to delete contact")
		return errors.New("failed to delete contact")
	}

	log.Info("Contact deleted successfully")
	return nil
}
