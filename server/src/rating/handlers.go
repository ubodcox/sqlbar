package rating

import (
	"net/http"
	"sqlbar/server/src/logs"
)

func init() {
	logs.Log.PushFuncName("users", "handlers", "init")
	defer logs.Log.PopFuncName()

	logs.Log.Info("IMPORTED")

	http.HandleFunc("/api/v1/rating/get", api.ratingGet)
}
