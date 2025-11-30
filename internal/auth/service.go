package auth

import (
	"context"
	"errors"
	"revaultier/configuration"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	cfg            *configuration.Config
	authRepository *AuthRepository
}

func NewAuthService(cfg *configuration.Config, authRepository *AuthRepository) *AuthService {
	return &AuthService{
		cfg:            cfg,
		authRepository: authRepository,
	}
}

func (s *AuthService) AuthenticateService(ctx context.Context, username string, password string) error {
	// Returns err if the authentication failed and nil if success

	exists, err := s.authRepository.UserExists(ctx, username)

	if err != nil {
		log.Error(err)
		return errors.New("could not retrieve user")
	}

	if !exists {
		return errors.New("user does not exist")
	}

	user, err := s.authRepository.GetUser(ctx, username)

	if err != nil {
		return errors.New("could not retrieve user")
	}

	if err := s.verifyPassword(user.Password, password); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignupService(ctx context.Context, username string, password string) error {
	exists, err := s.authRepository.UserExists(ctx, username)

	if err != nil {
		log.Error(err)
		return errors.New("could not retrieve user")
	}

	if exists {
		return ErrUsernameTaken
	}

	userId := uuid.New().String()

	if userId == "" {
		return errors.New("could not generate uuid")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return errors.New("could not generate hash from password")
	}

	if err := s.authRepository.InserUser(ctx, userId, username, string(passwordHash)); err != nil {
		return errors.New("could not add user")
	}

	return nil
}

func (s *AuthService) GenerateTokenSerivce(username string) (string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(s.cfg.Auth.SecretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) verifyPassword(passwordHash string, password string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return err
	}

	return nil
}
