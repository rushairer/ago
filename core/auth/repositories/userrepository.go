package repositories

import (
	"errors"
	"time"

	"PACKAGENAME/core/auth/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidUserDetail = errors.New("invalid user detail")
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	sqlString := `
	select
		id,
		name,
		email,
		connected_account_id,
		created_at,
		updated_at,
		email_verified_at
	from
		users
	where
		email = ?
	`
	if err := r.db.Get(user, sqlString, email); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (r *UserRepository) FindUserById(id string) (*models.User, error) {
	user := &models.User{}
	sqlString := `
	select
		id,
		name,
		email,
		connected_account_id,
		created_at,
		updated_at,
		email_verified_at
	from
		users
	where
		id = ?
	`
	if err := r.db.Get(user, sqlString, id); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (r *UserRepository) SaveUser(user *models.User) error {
	now := time.Now()

	if len(user.Id) == 0 {
		user.Id = uuid.NewString()
		user.CreatedAt = now
	}
	user.UpdatedAt = now

	if len(user.Email) == 0 {
		return ErrInvalidEmail
	} else {
		if len(user.Name) == 0 {
			user.Name = user.Email
		}
	}

	_, err := r.db.NamedExec(`
	replace into users
		(id, name, email, connected_account_id, created_at, updated_at)
	values
		(:id, :name, :email, :connected_account_id, :created_at, :updated_at)
	`, user)

	return err
}
