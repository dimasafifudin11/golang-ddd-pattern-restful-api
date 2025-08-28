package application

import (
	"context"
	"errors"
	"time"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/common"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/repository"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const dbTimeout = 2 * time.Second

type userServiceImpl struct {
	userRepository repository.UserRepository
	log            *logrus.Logger
}

func NewUserService(userRepo repository.UserRepository, log *logrus.Logger) service.UserService {
	return &userServiceImpl{
		userRepository: userRepo,
		log:            log,
	}
}

func (s *userServiceImpl) GetProfile(ctx context.Context, userID uint) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	// ... (rest of the implementation is the same as before)
	log := s.log.WithFields(logrus.Fields{"method": "GetProfile", "userID": userID})
	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn("User profile not found")
			return nil, errors.New("user not found")
		}
		log.WithError(err).Error("Failed to find user by ID")
		return nil, errors.New("server error")
	}
	log.Info("Successfully retrieved user profile")
	return user, nil
}

func (s *userServiceImpl) GetAllUsers(ctx context.Context) ([]model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()
	// ... (rest of the implementation is the same as before)
	log := s.log.WithField("method", "GetAllUsers")
	users, err := s.userRepository.FindAll(ctx)
	if err != nil {
		log.WithError(err).Error("Failed to get all users")
		return nil, errors.New("server error")
	}
	log.Info("Successfully retrieved all users")
	return users, nil
}

func (s *userServiceImpl) GetUserByID(ctx context.Context, userID uint) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "GetUserByID", "userID": userID})

	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn("User not found")
			return nil, common.ErrNotFound // <-- Mengembalikan error yang benar
		}
		log.WithError(err).Error("Failed to find user by ID")
		return nil, errors.New("server error")
	}

	log.Info("Successfully retrieved user by ID")
	return user, nil
}

func (s *userServiceImpl) UpdateUser(ctx context.Context, userID uint, request service.UserUpdateRequest) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "UpdateUser", "userID": userID})

	user, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn("User not found for update")
			return nil, common.ErrNotFound // <-- Mengembalikan error yang benar
		}
		log.WithError(err).Error("Failed to find user for update")
		return nil, errors.New("server error")
	}

	user.Name = request.Name
	user.Email = request.Email

	if err := s.userRepository.Update(ctx, user); err != nil {
		log.WithError(err).Error("Failed to update user")
		return nil, errors.New("failed to update user")
	}

	log.Info("User updated successfully")
	return user, nil
}

func (s *userServiceImpl) DeleteUser(ctx context.Context, userID uint) error {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	log := s.log.WithFields(logrus.Fields{"method": "DeleteUser", "userID": userID})

	_, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn("User not found for deletion")
			return common.ErrNotFound // <-- Mengembalikan error yang benar
		}
		log.WithError(err).Error("Failed to find user for deletion")
		return errors.New("server error")
	}

	if err := s.userRepository.Delete(ctx, userID); err != nil {
		log.WithError(err).Error("Failed to delete user")
		return errors.New("failed to delete user")
	}

	log.Info("User deleted successfully")
	return nil
}
