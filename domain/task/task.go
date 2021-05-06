package task

type Task struct {
	ID     string `datastore:"-"`
	Title  string
	Detail string
}
