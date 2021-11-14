package domain

import "time"

type User struct {
	ID        uint
	Name      string
	Password  string
	FirstName string
	LastName  string
	BirthDate time.Time
	GenderID  uint
	timeStampWithDelete
}

type Gender struct {
	ID      uint
	NameRus string
	NameEng string
}
