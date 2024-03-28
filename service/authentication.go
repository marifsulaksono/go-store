package service

import (
	"context"
	"gostore/entity"
	"gostore/repo"
	"gostore/utils/helper"
	authError "gostore/utils/helper/domain/errorModel"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authenticationService struct {
	Repo     repo.AuthentiocationRepo
	UserRepo repo.UserRepository
}

type AuthenticationService interface {
	LoginService(ctx context.Context, c *entity.Credential) (string, string, error)
	LogoutService(ctx context.Context, s *entity.Session) error
	RenewAccessToken(ctx context.Context, s *entity.Session) (string, error)
}

func NewAuthenticationService(r repo.AuthentiocationRepo, u repo.UserRepository) AuthenticationService {
	return &authenticationService{
		Repo:     r,
		UserRepo: u,
	}
}

func (a *authenticationService) LoginService(ctx context.Context, c *entity.Credential) (string, string, error) {
	user, err := a.UserRepo.GetUser(ctx, 0, c.Username)
	if err != nil {
		return "", "", authError.ErrLogin
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.Password))
	if err != nil {
		return "", "", authError.ErrLogin
	}

	accessTokenExpTime := time.Now().Add(time.Minute * 30)
	accessToken, err := helper.GenerateToken(user, accessTokenExpTime)
	if err != nil {
		return "", "", err
	}

	refreshTokenExpTime := time.Now().Add(time.Hour * 24 * 30)
	refreshToken, err := helper.GenerateToken(user, refreshTokenExpTime)
	if err != nil {
		return "", "", err
	}

	session := entity.Session{
		Id:    uuid.New(),
		Token: refreshToken,
	}

	if err := a.Repo.CreateSession(ctx, &session); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (a *authenticationService) LogoutService(ctx context.Context, s *entity.Session) error {
	return a.Repo.Logout(ctx, s.Token)
}

func (a *authenticationService) RenewAccessToken(ctx context.Context, s *entity.Session) (string, error) {
	if err := a.Repo.ValidateRefreshToken(ctx, s.Token); err != nil {
		return "", authError.ErrInvalidToken
	}

	payload, err := helper.VerifyRefreshToken(*s)
	if err != nil {
		return "", err
	}

	user := entity.UserResponse{
		Id:       payload.Id,
		Username: payload.Username,
		Email:    payload.Email,
	}

	accessTokenExpTime := time.Now().Add(time.Minute * 30)
	return helper.GenerateToken(user, accessTokenExpTime)
}
