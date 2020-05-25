package rating

import (
	"encoding/json"
	"net/http"
	"sqlbar/server/src/core"
	"sqlbar/server/src/logs"
)

type (
	// API struct
	API struct {
	}
)

var api API

func (api *API) ratingGet(w http.ResponseWriter, r *http.Request) {
	logs.Log.PushFuncName("rating", "api", "ratingGet")
	defer logs.Log.PopFuncName()

	err := r.ParseForm()
	if err != nil {
		logs.Log.Error("ParseForm", err)
		//json.NewEncoder(w).Encode(core.MakeDefaultResponse(-1, "error parse form", session))
		return
	}

	mode := core.ParseStr(r, "mode", "")
	dbRating, err := db.ratingGet(mode)
	if err != nil {
		logs.Log.Error("ParseForm", err)
		//json.NewEncoder(w).Encode(core.MakeDefaultResponse(-1, "error parse form", session))
		return
	}

	webRating := dbRating.MakeWeb()
	data, err := json.Marshal(webRating)
	if err != nil {
		logs.Log.Error("Marshal", err)
		//json.NewEncoder(w).Encode(core.MakeDefaultResponse(-1, "error Marshal", err, session))
		//http.Error(w, err.Error(), 500)
		return
	}

	w.Write(data)
}
