package secrets

import "time"

type Secret struct {
	Id                   string
	ExpiryTime           time.Time
	MaxRetrievalCount    int
	RetrievalCount       int
	EncryptedSecret      string
	InitializationVector string
}
