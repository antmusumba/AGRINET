package services

import (
	"errors"

	"github.com/antmusumba/agrinet/internals/models"
	"github.com/antmusumba/agrinet/internals/repositories"
	"github.com/antmusumba/agrinet/pkg"
)

type AuthService struct {
	repo repositories.UserRepo
}

func NewAuthService(repo repositories.UserRepo) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(user *models.User) error {
	existingUser, _ := s.repo.GetUserByEmail(user.Email)
	if existingUser != nil {
		return errors.New("this email is already in use")
	}

	hashedPassword, _ := pkg.HashPassword(user.Password)
	user.Password = hashedPassword
	return s.repo.CreateUser(user)
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("check your credentials and try again")
	}

	err = pkg.CheckPassword(user.Password, password)
	if err != nil {
		return nil, errors.New("invalid username or password")
	}
	return user, nil
}
