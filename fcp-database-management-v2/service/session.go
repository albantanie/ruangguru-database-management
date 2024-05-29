package service

import (
	"a21hc3NpZ25tZW50/model"
	"a21hc3NpZ25tZW50/repository"
	"fmt"
	"time"
)

type SessionService interface {
	AddSession(session model.Session) error
	UpdateSession(session model.Session) error
	DeleteSession(sessionToken string) error
	SessionAvailName(username string) error
	TokenExpired(session model.Session) bool
	TokenValidity(token string) (model.Session, error)
}

type sessionService struct {
	sessionRepository repository.SessionsRepository
}

func NewSessionService(sessionRepository repository.SessionsRepository) SessionService {
	return &sessionService{sessionRepository}
}

func (s *sessionService) SessionAvailName(username string) error {
	return s.sessionRepository.SessionAvailName(username)
}

func (s *sessionService) AddSession(session model.Session) error {
	return s.sessionRepository.AddSessions(session)
}

func (s *sessionService) UpdateSession(session model.Session) error {
	return s.sessionRepository.UpdateSessions(session)
}

func (s *sessionService) DeleteSession(sessionToken string) error {
	return s.sessionRepository.DeleteSession(sessionToken)
}

func (s *sessionService) TokenValidity(token string) (model.Session, error) {
	session, err := s.sessionRepository.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if s.TokenExpired(session) {
		err := s.sessionRepository.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}

	return session, nil
}

func (s *sessionService) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
