package application

import (
	"context"
	"errors"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/repository"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

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
