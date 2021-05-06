package datastore

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"

	"github.com/qushot/datastore-with-goroutine/domain/user"
)

const userKind = "User"

// NewUserRepository implements user.Repository interface.
func NewUserRepository() user.Repository {
	return &userRepository{}
}

type userRepository struct{}

func (r *userRepository) Save(ctx context.Context, u *user.User) (*user.User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	key := datastore.NameKey(userKind, id.String(), nil)
	key.Namespace = namespace

	if _, err := client.Put(ctx, key, u); err != nil {
		return nil, err
	}

	u.ID = id.String()

	return u, nil
}

func (r *userRepository) FetchAll(ctx context.Context) ([]*user.User, error) {
	q := datastore.NewQuery(userKind).Namespace(namespace)
	var us []*user.User
	keys, err := client.GetAll(ctx, q, &us)
	if err != nil {
		return nil, err
	}

	for i := range us {
		us[i].ID = keys[i].Name
	}

	return us, nil
}

func (r *userRepository) DeleteAll(ctx context.Context) error {
	q := datastore.NewQuery(userKind).Namespace(namespace).KeysOnly()
	keys, err := client.GetAll(ctx, q, nil)
	if err != nil {
		return err
	}

	if err := client.DeleteMulti(ctx, keys); err != nil {
		return err
	}

	return nil
}
