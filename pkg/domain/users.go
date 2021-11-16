package domain

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Password  string
	FirstName string
	LastName  string
	BirthDate time.Time
	GenderID  uint
	RoleID  uint
	timeStampWithDelete
}

type Role struct {
	ID uuid.UUID
	Name string
	Description string
}

type Gender struct {
	ID      uint
	NameRus string
	NameEng string
}
