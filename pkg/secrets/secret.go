package secrets

import "time"

type Secret struct {
	Id                      string
	DisableExpiryTime       bool
	ExpiryTime              time.Time
	AllowUnlimitedRetrieval bool
	MaxRetrievalCount       int
	RetrievalCount          int
	EncryptedSecret         string
	InitializationVector    string
	CreatedAt               time.Time
}
