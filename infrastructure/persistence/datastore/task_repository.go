package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"

	"github.com/qushot/datastore-with-goroutine/domain/task"
)

const taskKind = "Task"

// NewTaskRepository implements task.Repository interface.
func NewTaskRepository() task.Repository {
	return &taskRepository{}
}

type taskRepository struct{}

func (r *taskRepository) Save(ctx context.Context, t *task.Task) (*task.Task, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	key := datastore.NameKey(taskKind, id.String(), nil)
	key.Namespace = namespace

	if _, err := client.Put(ctx, key, t); err != nil {
		return nil, err
	}

	t.ID = id.String()

	return t, nil
}

func (r *taskRepository) FetchByID(ctx context.Context, id string) (*task.Task, error) {
	key := datastore.NameKey(taskKind, id, nil)
	key.Namespace = namespace

	t := &task.Task{}
	if err := client.Get(ctx, key, t); err != nil {
		return nil, err
	}

	t.ID = id

	return t, nil
}

func (r *taskRepository) DeleteAll(ctx context.Context) error {
	q := datastore.NewQuery(taskKind).Namespace(namespace).KeysOnly()
	keys, err := client.GetAll(ctx, q, nil)
	if err != nil {
		return err
	}

	if err := client.DeleteMulti(ctx, keys); err != nil {
		return err
	}

	return nil
}
