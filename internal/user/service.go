package user

import (
	"errors"

	"github.com/YasenMakioui/revaultier/internal/utils"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// import (
// 	"errors"

// 	"github.com/YasenMakioui/revaultier/internal/utils"
// )

type UserService struct {
	userRepository *UserRepository
}

func NewService(ur *UserRepository) *UserService {
	return &UserService{userRepository: ur}
}

func (s *UserService) UserExists(email string) bool {
	return s.userRepository.UserExists(email)
}

func (s *UserService) Signup(email string, password string) error {

	hashedPassword, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	if err := s.userRepository.InsertUser(email, hashedPassword); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Authenticate(username string, password string) error {
	// return error if the authentication fails in one of the steps

	if s.userRepository.UserExists(username) {
		return errors.New("user already exists")
	}

	return nil
}

func (s *UserService) CreateSession(c echo.Context, username string) error {
	sess, err := session.Get("session", c)

	if err != nil {
		return err
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	sess.Values["username"] = username
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}

// func (s *Service) Authenticate(email, password string) error {

// 	users := map[string]string{
// 		"user@example.com": "hashedpassword",
// 	}

// 	// check if user exists

// 	if email != "user@example.com" || !utils.CheckPasswordHash(password, users[email]) {
// 		return errors.New("invalid credentials")
// 	}

// 	return nil
// }

// func (s *Service) CreateSession() error {
// 	return nil
// }
