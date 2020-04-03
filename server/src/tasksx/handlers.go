package tasksx

import (
	"net/http"
	"sqlbar/server/src/core"
	"sqlbar/server/src/logs"
	//"time"
)

func init() {
	logs.Log.PushFuncName("tasks", "handlers", "init")
	defer logs.Log.PopFuncName()

	logs.Log.Info("IMPORTED")

	//chainMiddleWare := core.ChainMiddleWare(core.WithLogging, core.WithAuth)

	//auth
	http.HandleFunc("/main.html", func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "tasks", "main")
	})

	// api
	//http.HandleFunc("/api/v1/tasks/test", api.test)
	http.HandleFunc("/api/v1/tasks/get", api.get)
	http.HandleFunc("/api/v1/tasks/check", api.check)
}
