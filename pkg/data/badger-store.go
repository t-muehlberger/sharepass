package data

import (
	"encoding/json"
	"fmt"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/t-muehlberger/sharepass/pkg/secrets"
)

func NewDiskStore(path string) (secrets.Store, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, fmt.Errorf("failed to open data store: %w", err)
	}

	return &badgerStore{
		db: db,
	}, nil
}

func NewMemoryStore() (secrets.Store, error) {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))

	if err != nil {
		return nil, fmt.Errorf("failed to create data store: %w", err)
	}

	return &badgerStore{
		db: db,
	}, nil
}

type badgerStore struct {
	db *badger.DB
}

func (b *badgerStore) Put(secret secrets.Secret) error {
	key := []byte(secret.Id)
	value, err := json.Marshal(secret)
	if err != nil {
		return err
	}

	err = b.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})

	return err
}

func (b *badgerStore) Get(id string) (secrets.Secret, error) {
	key := []byte(id)
	var value []byte
	var s secrets.Secret

	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == badger.ErrKeyNotFound {
			return secrets.NotFoundError
		} else if err != nil {
			return err
		}

		value, err = item.ValueCopy(value)
		return err
	})
	if err != nil {
		return s, err
	}

	err = json.Unmarshal(value, &s)
	return s, err
}

func (b *badgerStore) Delete(id string) error {
	key := []byte(id)

	err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Delete(key)
	})

	if err != nil {
		return err
	}

	return nil
}

func (b *badgerStore) IncrementRetreivalCount(id string) (secrets.Secret, error) {
	key := []byte(id)
	var s secrets.Secret

	err := b.db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == badger.ErrKeyNotFound {
			return secrets.NotFoundError
		} else if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			return json.Unmarshal(val, &s)
		})

		s.RetrievalCount++

		value, err := json.Marshal(s)
		if err != nil {
			return err
		}

		return txn.Set(key, value)
	})

	return s, err
}

func (b *badgerStore) Count() (int, error) {
	var count int
	err := b.db.View(func(txn *badger.Txn) error {
		// iterator is the only way to count in badger
		iter := txn.NewIterator(badger.IteratorOptions{})
		defer iter.Close()
		for iter.Rewind(); iter.Valid(); iter.Next() { // https://dgraph.io/docs/badger/get-started/#iterating-over-keys
			count++
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (b *badgerStore) Close() error {
	return b.db.Close()
}
