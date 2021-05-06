package handler

import (
	"fmt"
	"net/http"

	"github.com/qushot/datastore-with-goroutine/di"
)

// TearDown cleans datastore in 'goroutinetest' namespace.
func TearDown(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	container := di.Get()

	if err := container.TaskRepository.DeleteAll(ctx); err != nil {
		http.Error(w, "taskRepository.DeleteAll error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := container.UserRepository.DeleteAll(ctx); err != nil {
		http.Error(w, "userRepository.DeleteAll error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "done")
}
