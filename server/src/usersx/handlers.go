package users

import (
	"net/http"
	"sqlbar/server/src/core"
	"sqlbar/server/src/logs"
	"time"
)

//TODO: move to core
//WriteCookie func
func WriteCookie(w http.ResponseWriter, r *http.Request, user User) {
	logs.Log.PushFuncName("users", "users", "WriteCookie")
	defer logs.Log.PopFuncName()

	token := core.Create(user.Email, user.Password)
	logs.Log.Debug("BEGIN token create:", token) //, "sessions:", core.Sessions.Cache)

	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := &http.Cookie{Name: "token", Value: token, Expires: expiration,
		Path: "/", MaxAge: 86400}
	http.SetCookie(w, cookie)

	var session *core.Session
	var err error
	if session, err = core.CheckAuth(token); err != nil {
		logs.Log.Error("core.CheckAuth('token') err:", err)
		http.Redirect(w, r, "../../auth.html", 303)
		return
	}

	session.Expired = expiration
	session.ID = user.ID
	session.Password = user.Password
	session.FirstName = user.FirstName
	session.LastName = user.LastName
	session.Post = user.Post
	session.Phone = user.Phone
	session.Email = user.Email
	session.Role = user.Role
	session.Registered = user.Registered
	core.UpdateSession(token, session)
	logs.Log.Debug("END", "session:", session) //, "sessions:", core.Sessions.Cache)
}

func init() {
	logs.Log.PushFuncName("users", "handlers", "init")
	defer logs.Log.PopFuncName()

	logs.Log.Info("IMPORTED")

	chainMiddleWare := core.ChainMiddleWare(core.WithLogging, core.WithAuth)

	//auth
	/* http.HandleFunc("/auth.html", func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "users", "auth")
	})
	// html
	http.HandleFunc("/users/data.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "users", "data")
	}))
	http.HandleFunc("/users/insert.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "users", "insert")
	}))
	http.HandleFunc("/users/update.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "users", "update")
	}))
	http.HandleFunc("/users/profile.html", chainMiddleWare(func(w http.ResponseWriter, r *http.Request) {
		core.DefaultHandler(w, r, "users", "profile")
	})) */

	/* http.HandleFunc("/users/exec_insert.html", usersExecInsertHandler)
	http.HandleFunc("/users/exec_update.html", usersExecUpdateHandler)
	http.HandleFunc("/users/exec_delete.html", usersExecDeleteHandler) */

	// api
	//http.HandleFunc("/api/v1/users/noauth", api.noauth)
	http.HandleFunc("/api/v1/users/auth", api.auth)
	http.HandleFunc("/api/v1/users/data", chainMiddleWare(api.data))
	http.HandleFunc("/api/v1/users/insert", chainMiddleWare(api.insert))
	http.HandleFunc("/api/v1/users/update", chainMiddleWare(api.update))
	http.HandleFunc("/api/v1/users/delete", chainMiddleWare(api.delete))
}
