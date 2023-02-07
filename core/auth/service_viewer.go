package auth

import (
	"PACKAGENAME/core/auth/models"
	"PACKAGENAME/core/auth/repositories"

	"github.com/jmoiron/sqlx"
)

type ViewerService struct {
	Viewer         *models.User
	userRepository *repositories.UserRepository
	db             *sqlx.DB
}

func NewViewerService(db *sqlx.DB) *ViewerService {
	userRepository := repositories.NewUserRepository(db)

	return &ViewerService{
		db:             db,
		userRepository: userRepository,
	}
}

func (s *ViewerService) FindViewer(id string) error {
	viewer, err := s.userRepository.FindUserById(id)
	s.Viewer = viewer
	return err
}
