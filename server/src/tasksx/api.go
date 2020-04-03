package tasksx

import (
	"encoding/json"
	"net/http"
	"sqlbar/server/src/core"
	"sqlbar/server/src/logs"
	"strconv"
)

type (
	// API struct
	API struct {
	}

	// response struct
	getTaskResponse struct {
		core.EmptyResponse
		Tasks []Task `json:"tasks"`
	}

	checkTaskResponse struct {
		core.EmptyResponse
		ID     int64      `json:"id"`
		Result ResultList `json:"result"`
	}
)

var api API

// makeResponse func
func makeGetTaskResponse(code int, msg string, s *core.Session, tasks []Task) *getTaskResponse {
	return &getTaskResponse{
		*core.MakeEmptyResponse(code, msg),
		tasks,
	}
}

func makeCheckTaskResponse(code int, msg string, s *core.Session, id int64, result ResultList) *checkTaskResponse {
	return &checkTaskResponse{
		*core.MakeEmptyResponse(code, msg),
		id, result,
	}
}

func (api *API) get(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("tasks", "api", "get")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("start")
	defer logs.Log.Debug("end")

	//json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "test"))

	core.EnableCors(&w)

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "error parse form"))
		return
	}

	var id int64 = -1
	if _, ok := r.Form["id"]; ok {
		id, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}
	/* if id < 1 {
		logs.Log.Error("no id param", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "no id param"))
		return
	} */

	//logs.Log.Debug("id:", id)

	tasks, err := sql.getTask(id)
	if err != nil {
		logs.Log.Error("sql.getTask", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, err.Error()))
		return
	}

	//logs.Log.Debug("test 2", id, text, cols, stars)

	json.NewEncoder(w).Encode(
		makeGetTaskResponse(1, "", nil, tasks),
	)

	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(core.MakeDefaultResponse(-1, "error parse form", nil))

	return
}

func (api *API) check(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("tasks", "api", "check")
	defer logs.Log.PopFuncName()

	logs.Log.Debug("start")
	defer logs.Log.Debug("end")

	core.EnableCors(&w)

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		json.NewEncoder(w).Encode(core.MakeDefaultResponse(-1, "error parse form", nil))
		return
	}

	var id int64 = -1
	if _, ok := r.Form["id"]; ok {
		id, _ = strconv.ParseInt(r.Form["id"][0], 10, 64)
	}
	if id < 1 {
		logs.Log.Error("no id", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "no id param"))
		return
	}

	var request string = ""
	if _, ok := r.Form["sql"]; ok {
		request = r.Form["sql"][0]
	}
	if request == "" {
		logs.Log.Error("no sql param", err)
		json.NewEncoder(w).Encode(core.MakeEmptyResponse(-1, "no sql param"))
		return
	}

	/*result, err := sql.CheckTask(id, request)
	if err != nil {
		logs.Log.Error("sql.getTask", err)
		json.NewEncoder(w).Encode(core.MakeDefaultResponse(-1, err.Error(), nil))
		return
	}

	json.NewEncoder(w).Encode(
		makeCheckTaskResponse(1, "error parse form", nil, id, result),
	)*/

}
