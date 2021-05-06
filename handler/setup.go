package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/qushot/datastore-with-goroutine/di"
	"github.com/qushot/datastore-with-goroutine/domain/task"
	"github.com/qushot/datastore-with-goroutine/domain/user"
)

// SetUp creates test data.
func SetUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	container := di.Get()

	queryParamNum := r.URL.Query().Get("num")
	if queryParamNum == "" {
		queryParamNum = "5"
	}

	num, err := strconv.Atoi(queryParamNum)
	if err != nil {
		http.Error(w, "strconv.Atoi error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	for i := 1; i <= num; i++ {
		s := strconv.Itoa(i)
		t := &task.Task{
			Title:  "テストタイトル" + s,
			Detail: "詳細テスト" + s,
		}

		newTask, err := container.TaskRepository.Save(ctx, t)
		if err != nil {
			http.Error(w, "TaskRepository.Save error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		u := &user.User{
			Name:   "テスト" + s + "さん",
			TaskID: newTask.ID,
		}

		if _, err := container.UserRepository.Save(ctx, u); err != nil {
			http.Error(w, "UserRepository.Save error: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	fmt.Fprint(w, "done")
}
