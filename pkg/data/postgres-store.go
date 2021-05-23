package data

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/t-muehlberger/sharepass/pkg/secrets"
)

type postgresStore struct {
	db *pg.DB
}

func NewPostgresStore(dbHost, dbUser, dbPassword, dbName string) (secrets.Store, error) {
	db := pg.Connect(&pg.Options{
		Addr:     dbHost,
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})

	err := createPgSchema(db)
	if err != nil {
		return nil, err
	}

	return &postgresStore{
		db: db,
	}, nil
}

func (s *postgresStore) Put(secret secrets.Secret) error {
	_, err := s.db.Model(&secret).Insert()

	return err
}

func (s *postgresStore) Get(id string) (secrets.Secret, error) {
	secret := secrets.Secret{Id: id}
	err := s.db.Model(&secret).WherePK().Select()

	return secret, err
}

func (s *postgresStore) Delete(id string) error {
	secret := secrets.Secret{Id: id}
	_, err := s.db.Model(&secret).WherePK().Delete()

	return err
}

func (s *postgresStore) IncrementRetreivalCount(id string) (secrets.Secret, error) {
	sec := secrets.Secret{Id: id}
	err := s.db.RunInTransaction(context.Background(), func(tx *pg.Tx) error {
		err := tx.Model(&sec).WherePK().First()
		if err != nil {
			return err
		}
		sec.RetrievalCount++
		_, err = tx.Model(&sec).WherePK().Update()
		return err
	})
	return sec, err
}

func (s *postgresStore) Close() error {
	return s.db.Close()
}

// createSchema creates database schema for User and Story models.
func createPgSchema(db *pg.DB) error {
	models := []interface{}{
		(*secrets.Secret)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
