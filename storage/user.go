package storage

import (
	"github.com/google/uuid"
)

type User struct {
	UserId        uuid.UUID
	FirebaseToken string
}
