package secrets

type constError string

func (err constError) Error() string {
	return string(err)
}

var NotFoundError error = constError("Secret not found")

type Store interface {
	Put(secret Secret) error
	Get(id string) (Secret, error)
	IncrementRetreivalCount(id string) (Secret, error)
	Delete(id string) error
	Close() error
}
