package data

import (
	"errors"
	"testing"

	"github.com/t-muehlberger/sharepass/pkg/secrets"
)

const secretId = "1234"   // id of existing item in the test store
const unknonwnId = "4321" // if of non existing item in the test store

func setupSecret(id string) secrets.Secret {
	return secrets.Secret{
		Id:                id,
		MaxRetrievalCount: 3,
	}
}

func setupStore() secrets.Store {
	store := NewMemoryStore()
	secret := setupSecret(secretId)
	store.Put(secret)
	return store
}

func TestStoreGet(t *testing.T) {
	var store = setupStore()

	// get non existing itme
	_, err := store.Get(unknonwnId)
	if !errors.Is(err, secrets.NotFoundError) {
		t.Fail()
	}

	// get existing item
	s, err := store.Get(secretId)
	if err != nil || s.Id != secretId {
		t.Fail()
	}
}

func TestStorePut(t *testing.T) {
	var store = setupStore()
	const putId = "test-put-id"
	var secret = setupSecret(putId)

	err := store.Put(secret)
	if err != nil {
		t.Fail()
	}

	// verify put worked
	s, err := store.Get(putId)
	if err != nil || s.Id != putId {
		t.Fail()
	}
}

func TestStoreDelete(t *testing.T) {
	var store = setupStore()

	// delete non existing object
	err := store.Delete(unknonwnId)
	if err != nil {
		t.Fail()
	}

	// delete existing object
	err = store.Delete(secretId)
	if err != nil {
		t.Fail()
	}

	// verify deleted
	_, err = store.Get(secretId)
	if !errors.Is(err, secrets.NotFoundError) {
		t.Fail()
	}
}

func TestIncrementRetreivalCount(t *testing.T) {
	var store = setupStore()

	// increment non existing item
	_, err := store.IncrementRetreivalCount(unknonwnId)
	if !errors.Is(err, secrets.NotFoundError) {
		t.Fail()
	}

	// increment existing
	s, err := store.IncrementRetreivalCount(secretId)
	if err != nil || s.RetrievalCount != 1 || s.Id != secretId {
		t.Fail()
	}
}
