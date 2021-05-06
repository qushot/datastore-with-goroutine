package user

type User struct {
	ID     string `datastore:"-"`
	Name   string
	TaskID string
}
