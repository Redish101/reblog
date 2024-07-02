package core

import "reblog/internal/auth"

type AuthService struct {
	app *App

	auth *auth.Auth
}

func NewAuthService(app *App) *AuthService {
	return &AuthService{
		app: app,
	}
}

func (s *AuthService) Start() error {
	s.auth = auth.NewAuth(s.app.Query())

	return nil
}

func (s *AuthService) Stop() error {
	s.auth = nil

	return nil
}

func (s *AuthService) GetToken(username string, password string) string {
	return s.auth.GetToken(username, password)
}

func (s *AuthService) VerifyToken(token string) bool {
	return s.auth.VerifyToken(token)
}
