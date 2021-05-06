package task

import "context"

type Repository interface {
	Save(ctx context.Context, t *Task) (*Task, error)
	FetchByID(ctx context.Context, id string) (*Task, error)
	DeleteAll(ctx context.Context) error
}
