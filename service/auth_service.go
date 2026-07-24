package service

import (
	"errors"

	"github.com/dewialvi/digital-channel-monitoring/config"
	"github.com/dewialvi/digital-channel-monitoring/models"
	"github.com/dewialvi/digital-channel-monitoring/repository"
	"github.com/dewialvi/digital-channel-monitoring/utils"
)

type AuthService struct {
	UserRepo *repository.UserRepository
	Cfg      *config.Config
}

func NewAuthService(
	userRepo *repository.UserRepository,
	cfg *config.Config,
) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
		Cfg:      cfg,
	}
}

var ErrInvalidCredentials = errors.New("email atau password salah")

var ErrAccountInactive = errors.New(
	"akun tidak aktif, hubungi admin",
)

func (s *AuthService) Register(
	name string,
	email string,
	password string,
	role models.Role,
) (*models.User, error) {

	existingUser, _ := s.UserRepo.FindByEmail(email)

	if existingUser != nil {
		return nil, errors.New("email sudah terdaftar")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
		Role:     role,
		IsActive: true,
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(
	email string,
	password string,
) (string, *models.User, error) {

	user, err := s.UserRepo.FindByEmail(email)

	if err != nil {
		return "", nil, ErrInvalidCredentials
	}

	if !user.IsActive {
		return "", nil, ErrAccountInactive
	}

	if !utils.CheckPasswordHash(
		password,
		user.Password,
	) {
		return "", nil, ErrInvalidCredentials
	}

	expiryHours := 24

	token, err := utils.GenerateJWT(
		user.ID,
		user.Email,
		string(user.Role),
		s.Cfg.JWTSecret,
		expiryHours,
	)

	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}