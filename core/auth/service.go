package auth

import (
	"errors"
	"net/http"

	"PACKAGENAME/core/auth/models"
	"PACKAGENAME/core/auth/repositories"

	"github.com/jmoiron/sqlx"
	"github.com/markbates/goth"
)

type AuthService struct {
	Auth             *models.User
	RedirectUrl      string
	AcceptOAuthScope bool
	userRepository   *repositories.UserRepository
	db               *sqlx.DB
}

var (
	ErrUserSessionIsEmpty        = errors.New("user session is empty")
	ErrRedirectUrlSessionIsEmpty = errors.New("redirect_url session is empty")
)

// 通过 goth.User 初始化 User
func NewUser(gothUser goth.User) *models.User {
	return &models.User{
		Name:  gothUser.Name,
		Email: gothUser.Email,
	}
}

func NewAuthService(db *sqlx.DB) *AuthService {
	user := NewUser(goth.User{})
	userRepository := repositories.NewUserRepository(db)

	return &AuthService{
		Auth:           user,
		db:             db,
		userRepository: userRepository,
	}
}

// 获取当前登录的用户信息。
func (s *AuthService) FindCurrentUser(r *http.Request) error {
	auth, err := s.userRepository.FindUserById(s.Auth.Id)
	s.Auth = auth
	return err
}
