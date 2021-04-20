package data

import (
	"sync"

	"github.com/t-muehlberger/sharepass/pkg/secrets"
)

func NewMemoryStore() secrets.Store {
	return &inMemory{
		mu:   sync.Mutex{},
		data: make(map[string]secrets.Secret),
	}
}

type inMemory struct {
	mu   sync.Mutex
	data map[string]secrets.Secret
}

func (m *inMemory) Put(secret secrets.Secret) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.data[secret.Id] = secret

	return nil
}

func (m *inMemory) Get(id string) (secrets.Secret, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	s, ok := m.data[id]
	if !ok {
		return s, secrets.NotFoundError
	}
	return s, nil
}

func (m *inMemory) Delete(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.data, id)

	return nil
}

func (m *inMemory) IncrementRetreivalCount(id string) (secrets.Secret, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	s, ok := m.data[id]
	if !ok {
		return s, secrets.NotFoundError
	}

	s.RetrievalCount++
	m.data[id] = s

	return s, nil
}

func (m *inMemory) Close() error {
	return nil
}
