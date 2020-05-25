package tasks

import (
	"net/http"
	"sqlbar/server/src/logs"
	//"time"
)

func init() {
	logs.Log.PushFuncName("tasks", "handlers", "init")
	defer logs.Log.PopFuncName()

	logs.Log.Info("IMPORTED")

	http.HandleFunc("/api/v1/tasks/get", api.get)
	http.HandleFunc("/api/v1/tasks/check", api.check)
}
