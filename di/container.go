package di

import (
	"github.com/qushot/datastore-with-goroutine/domain/task"
	"github.com/qushot/datastore-with-goroutine/domain/user"
	"github.com/qushot/datastore-with-goroutine/infrastructure/persistence/datastore"
)

var con *container

func init() {
	con = &container{
		UserRepository: datastore.NewUserRepository(),
		TaskRepository: datastore.NewTaskRepository(),
	}
}

func Get() *container {
	return con
}

type container struct {
	UserRepository user.Repository
	TaskRepository task.Repository
}
