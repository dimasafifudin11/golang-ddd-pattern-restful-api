package application

import (
	"context"
	"errors"

	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/model"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/repository"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/internal/domain/service"
	"github.com/dimasafifudin11/golang-ddd-pattern-restful-api/pkg/util"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type authServiceImpl struct {
	userRepository repository.UserRepository
	log            *logrus.Logger
}

func NewAuthService(userRepo repository.UserRepository, log *logrus.Logger) service.AuthService {
	return &authServiceImpl{
		userRepository: userRepo,
		log:            log,
	}
}

func (s *authServiceImpl) Register(ctx context.Context, name, email, password string) (*model.User, error) {
	log := s.log.WithFields(logrus.Fields{"method": "Register", "email": email})

	// Check if user already exists
	existingUser, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.WithError(err).Error("Failed to check existing user")
		return nil, errors.New("server error")
	}
	if existingUser != nil {
		log.Warn("User already exists")
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		log.WithError(err).Error("Failed to hash password")
		return nil, errors.New("server error")
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	if err := s.userRepository.Save(ctx, user); err != nil {
		log.WithError(err).Error("Failed to save user")
		return nil, errors.New("failed to register user")
	}

	log.Info("User registered successfully")
	return user, nil
}

func (s *authServiceImpl) Login(ctx context.Context, email, password string) (string, error) {
	log := s.log.WithFields(logrus.Fields{"method": "Login", "email": email})

	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warn("User not found")
			return "", errors.New("invalid email or password")
		}
		log.WithError(err).Error("Failed to find user by email")
		return "", errors.New("server error")
	}

	if !util.CheckPasswordHash(password, user.Password) {
		log.Warn("Invalid password")
		return "", errors.New("invalid email or password")
	}

	token, err := util.GenerateToken(user.ID)
	if err != nil {
		log.WithError(err).Error("Failed to generate token")
		return "", errors.New("failed to login")
	}

	log.Info("User logged in successfully")
	return token, nil
}
