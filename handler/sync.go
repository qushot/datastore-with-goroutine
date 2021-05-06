package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qushot/datastore-with-goroutine/di"
	"github.com/qushot/datastore-with-goroutine/domain"
)

func Sync(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	container := di.Get()

	users, err := container.UserRepository.FetchAll(ctx)
	if err != nil {
		http.Error(w, "userRepository.FetchAll error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	userTasks := make([]*domain.UserTask, len(users))
	for i := range userTasks {
		t, err := container.TaskRepository.FetchByID(ctx, users[i].TaskID)
		if err != nil {
			http.Error(w, "taskRepository.FetchByID error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		userTasks[i] = &domain.UserTask{
			UserID:     users[i].ID,
			UserName:   users[i].Name,
			TaskID:     t.ID,
			TaskTitle:  t.Title,
			TaskDetail: t.Detail,
		}
	}

	bytes, err := json.Marshal(userTasks)
	if err != nil {
		http.Error(w, "json.Marshal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(bytes))
}
