package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/qushot/datastore-with-goroutine/di"
	"github.com/qushot/datastore-with-goroutine/domain"
)

func Async(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	container := di.Get()

	users, err := container.UserRepository.FetchAll(ctx)
	if err != nil {
		http.Error(w, "userRepository.FetchAll error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var wg sync.WaitGroup
	userTasks := make([]*domain.UserTask, len(users))
	for i := range userTasks {
		wg.Add(1)
		go func(i int, id string) {
			t, err := container.TaskRepository.FetchByID(ctx, id)
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
			wg.Done()
		}(i, users[i].TaskID)
	}

	wg.Wait()

	bytes, err := json.Marshal(userTasks)
	if err != nil {
		http.Error(w, "json.Marshal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(bytes))
}
