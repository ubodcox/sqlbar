package users

import (
	"net/http"
	"sqlbar/server/src/core"
	"sqlbar/server/src/logs"
	"time"
)

//WriteCookie func
func WriteCookie(w http.ResponseWriter, r *http.Request, user User) {
	logs.Log.PushFuncName("users", "users", "WriteCookie")
	defer logs.Log.PopFuncName()

	token := core.Create(user.Email, user.Password)
	logs.Log.Debug("BEGIN token create:", token) //, "sessions:", core.Sessions.Cache)
	//defer logs.Log.Debug("END", "session:", session) //, "sessions:", core.Sessions.Cache)

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

}

func init() {
	logs.Log.PushFuncName("users", "handlers", "init")
	defer logs.Log.PopFuncName()

	logs.Log.Info("IMPORTED")

	//chainMiddleWare := core.ChainMiddleWare(core.WithLogging, core.WithAuth)

	// api
	//http.HandleFunc("/api/v1/users/noauth", api.noauth)
	http.HandleFunc("/api/v1/users/signup", api.insert)
	http.HandleFunc("/api/v1/users/signin", api.auth)
	http.HandleFunc("/api/v1/users/signout", core.WithAuth(api.auth))

	http.HandleFunc("/api/v1/users/get", api.get)
	http.HandleFunc("/api/v1/users/set", core.WithAuth(api.set))
}
