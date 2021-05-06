package user

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, u *User) (*User, error)
	FetchAll(ctx context.Context) ([]*User, error)
	DeleteAll(ctx context.Context) error
}
