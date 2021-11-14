package domain

import "time"

type timeStamp struct {
	CreatedAt time.Time
	UpdateAt  time.Time
}

type timeStampWithDelete struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}
