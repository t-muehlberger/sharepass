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
	ttl := time.Duration(timeToLive) * time.Second
	expiryTime := time.Now().Add(ttl)

	if timeToLive < -1 {
		return Secret{}, fmt.Errorf("argument error, timeToLive is invalid")
	}

	if maxRetrievalCount < -1 {
		return Secret{}, fmt.Errorf("argument error, maxRetrievalCount is invalid")
	}

	sec := Secret{
		Id:                      uuid.NewString(),
		DisableExpiryTime:       timeToLive == -1,
		ExpiryTime:              expiryTime,
		AllowUnlimitedRetrieval: maxRetrievalCount == -1,
		RetrievalCount:          0,
		MaxRetrievalCount:       maxRetrievalCount,
		EncryptedSecret:         encryptedSecret,
		InitializationVector:    initializationVector,
		CreatedAt:               time.Now(),
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
	timeout := !s.ExpiryTime.After(time.Now()) && !s.DisableExpiryTime
	retievalCount := s.RetrievalCount >= s.MaxRetrievalCount && !s.AllowUnlimitedRetrieval
	return timeout || retievalCount
}
