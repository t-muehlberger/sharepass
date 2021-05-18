package secrets

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Store Store
}

func (s *Service) CreateSecret(encryptedSecret string, initializationVector string, timeToLive int, maxRetrievalCount int) (Secret, error) {
	if timeToLive <= 0 {
		return Secret{}, fmt.Errorf("argument error, timeToLive cannot be less than zero")
	}
	if maxRetrievalCount < 0 {
		return Secret{}, fmt.Errorf("argument error, maxRetrievalCount cannot be less than zero")
	}

	ttl := time.Duration(timeToLive) * time.Second
	sec := Secret{
		Id:                   uuid.NewString(),
		ExpiryTime:           time.Now().Add(ttl),
		RetrievalCount:       0,
		MaxRetrievalCount:    maxRetrievalCount,
		EncryptedSecret:      encryptedSecret,
		InitializationVector: initializationVector,
	}

	err := s.Store.Put(sec)
	if err != nil {
		return sec, fmt.Errorf("failed to create secret: %w", err)
	}

	return sec, nil
}

func (s *Service) GetSecretMetadata(id string) (Secret, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return Secret{}, fmt.Errorf("argument error, could not parse id: %w", err)
	}

	sec, err := s.Store.Get(id)
	if err != nil {
		return Secret{}, fmt.Errorf("failed to load secret metadata: %w", err)
	}

	if isExpired(sec) {
		return Secret{}, fmt.Errorf("secret is expired")
	}

	return sec, nil
}

func (s *Service) RevealSecret(id string) (Secret, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return Secret{}, fmt.Errorf("argument error, could not parse id: %w", err)
	}

	sec, err := s.Store.IncrementRetreivalCount(id)
	if err != nil {
		return Secret{}, fmt.Errorf("failed to load secret metadata: %w", err)
	}

	tmp := sec
	tmp.RetrievalCount-- // TODO: improve this
	if isExpired(tmp) {
		return Secret{}, fmt.Errorf("secret is expired")
	}

	return sec, nil
}

func isExpired(s Secret) bool {
	timeout := !s.ExpiryTime.After(time.Now())
	retievalCount := s.RetrievalCount >= s.MaxRetrievalCount
	return timeout || retievalCount
}
